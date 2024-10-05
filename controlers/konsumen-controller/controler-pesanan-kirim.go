package konsumencontrollers

import (
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"
	modelkonsumens "backend-jona-golang/models/model-konsumen"
	"backend-jona-golang/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePesananJasaKirim(ctx *gin.Context) {
	var input modelkonsumens.InputPesananJasaKirimKonsumen
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

	if err := databases.DB.Table("sub_categories").Where("id = ?", input.JasaId).First(&dataSubCategory).Error; err != nil {
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
	payloadBank.BankTransfer.Bank = dataBank.Type
	payloadBank.CustomExpiry.OrderTime = time.Now().In(time.FixedZone("WIB", 7*60*60)).Format("2006-01-02 15:04:05 +0700")
	payloadBank.CustomExpiry.ExpiryDuration = 30
	payloadBank.CustomExpiry.Unit = "minutes"

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
	dataPesanan.JasaId = uint64(dataSubCategory.ID)
	dataPesanan.CodePesanan = orderID
	dataPesanan.Status = "menunggu"
	dataPesanan.TransactionMidtrans = response.TransactionID
	dataPesanan.VaBank = response.VANumbers[0].VANumber
	dataPesanan.IdAddress = uint64(address.ID)
	dataPesanan.AlamatTujuan = input.IdAlamatTujuan

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
		OrderID:       orderID,
	}
	if err := databases.DB.Table("notifikasi_pembayarans").Create(&newNotification).Error; err != nil {
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
