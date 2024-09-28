package konsumencontrollers

import (
	"backend-jona-golang/config"
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"
	modelkonsumens "backend-jona-golang/models/model-konsumen"
	"backend-jona-golang/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListPesananKonsumen(ctx *gin.Context) {
	var data []modelkonsumens.PesananKonsumen
	userId, _ := ctx.Get("userID")

	// Ambil parameter limit dan page dari query string, atau gunakan default jika tidak ada
	limit := ctx.DefaultQuery("limit", "10")
	page := ctx.DefaultQuery("page", "1")

	// Konversi limit dan page menjadi integer
	limitInt, _ := strconv.Atoi(limit)
	pageInt, _ := strconv.Atoi(page)

	// Hitung offset berdasarkan limit dan page
	offset := (pageInt - 1) * limitInt

	err := databases.DB.Table("pesanan_konsumens").
		Where("user_id = ?", userId).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID, email, role, nama, profile")
		}).
		Preload("Bank").
		Preload("Jasa", func(db *gorm.DB) *gorm.DB {
			return db.Preload("CaegoryUtama", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Fitur")
			})
		}).
		Preload("Address").
		Limit(limitInt). // Terapkan limit
		Offset(offset).  // Terapkan offset
		Find(&data).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Pesanan not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve Pesanan",
			})

			return
		}
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Detail pesanan ditemukan",
		"data":    data,
	})
}

func DetailPesananKonsumen(ctx *gin.Context) {
	var data modelkonsumens.PesananKonsumen
	userId, _ := ctx.Get("userID")

	err := databases.DB.Table("pesanan_konsumens").
		Where("id = ? AND user_id = ?", ctx.Param("id"), userId).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID, email, role, nama, profile")
		}).
		Preload("Bank").
		Preload("Jasa", func(db *gorm.DB) *gorm.DB {
			return db.Preload("CaegoryUtama", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Fitur")
			})
		}).
		Preload("Address").
		First(&data).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Pesanan not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve Pesanan",
			})

			return
		}
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Detail pesanan ditemukan",
		"data":    data,
	})
}

func CreatePesanan(ctx *gin.Context) {
	var input modelkonsumens.InputPesananKonsumen
	var dataPesanan modelkonsumens.PesananKonsumen
	var dataUser models.Users
	var dataSubCategory models.SubCategory
	var dataCategory models.CaegoryUtama
	var dataBank models.DaftarBank
	var address models.Address
	var payloadBank utils.BankTransferPayload

	user_id, _ := ctx.Get("userID")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	if err := databases.DB.Table("addresses").Where("id = ?", input.IdAddress).First(&address).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Address not found",
		})
		return
	}

	if err := databases.DB.Table("users").Where("id = ?", user_id).First(&dataUser).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "User not found",
		})
		return
	}

	if err := databases.DB.Table("sub_categories").Where("id = ?", input.JasaBersiId).First(&dataSubCategory).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Sub Categories not found",
		})
		return
	}

	if err := databases.DB.Table("caegory_utamas").Where("id = ?", dataSubCategory.IdCategory).First(&dataCategory).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Caegory Utamas not found",
		})
		return
	}

	if err := databases.DB.Table("daftar_banks").Where("id = ?", input.MetodePembayaran).First(&dataBank).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Bank not found",
		})
		return
	}

	// Menghasilkan OrderID
	orderID := utils.GenerateOrderID("JONA")

	payloadBank.PaymentType = "bank_transfer"
	payloadBank.TransactionDetails.GrossAmount = dataSubCategory.Harga
	payloadBank.TransactionDetails.OrderID = orderID
	payloadBank.TransactionDetails.ExpireTime = time.Now().Add(30 * time.Minute).Format(time.RFC3339) // Set kadaluarsa 30 menit
	payloadBank.BankTransfer.Bank = dataBank.Nama

	response, err := utils.VaNumberBank(payloadBank)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to create bank transfer: " + err.Error(),
		})
		return
	}

	// Simpan data pesanan ke dalam database
	dataPesanan.UserID = uint64(dataUser.ID)
	dataPesanan.MetodePembayaran = uint64(dataBank.ID)
	dataPesanan.JasaBersiId = uint64(dataSubCategory.ID)
	dataPesanan.CodePesanan = orderID
	dataPesanan.Status = "menunggu"
	dataPesanan.TransactionMidtrans = response.TransactionID
	dataPesanan.VaBank = response.VANumbers[0].VANumber
	dataPesanan.IdAddress = uint64(address.ID)

	if err := databases.DB.Table("pesanan_konsumens").Create(&dataPesanan).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to save order: " + err.Error(),
		})
		return
	}

	// Create a new notification record
	newNotification := modelkonsumens.NotifikasiPembayaran{
		StatusPesanan: modelkonsumens.NotifikasiMenunggu,
		Description:   "Waktu Pembayaran 30:00",
		TransactionID: response.TransactionID,
		UserId:        uint64(dataUser.ID),
	}
	if err := databases.DB.Create(&newNotification).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to create notification",
		})
		return
	}

	// Mengirimkan respons berhasil kepada pengguna
	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "Order created successfully",
		"data":    dataPesanan,
	})
}

func NotifikasiPembayaran(ctx *gin.Context) {
	var notifikasi modelkonsumens.MidtransNotification
	// var pesanan modelkonsumens.PesananKonsumen

	// Bind JSON body to the NotifikasiPembayaran struct
	if err := ctx.ShouldBindJSON(&notifikasi); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Menghasilkan signature yang diharapkan dari data notifikasi
	expectedSignature := utils.GenerateMidtransSignature(notifikasi.OrderID, notifikasi.StatusCode, notifikasi.GrossAmount, config.SERVER_KEY_MIDTRANS)

	// Validasi signature
	if notifikasi.SignatureKey != expectedSignature {
		ctx.JSON(401, gin.H{
			"error":   true,
			"message": "Invalid signature. Notification rejected.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   true,
		"message": "Notification processed successfully",
	})
}
