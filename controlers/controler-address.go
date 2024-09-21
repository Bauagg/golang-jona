package controlers

import (
	"backend-jona-golang/databases"
	"backend-jona-golang/models"
	"regexp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAddress(ctx *gin.Context) {
	var addresses []models.Address
	userId, _ := ctx.Get("userID")

	err := databases.DB.Table("addresses").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID, email, role")
		}).
		Where("addresses.user_id = ?", userId).
		Find(&addresses).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to retrieve addresses",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Get data Address success",
		"datas":   addresses,
	})
}

func DetailAddress(ctx *gin.Context) {
	var addresses models.Address
	userId, _ := ctx.Get("userID")
	addressID := ctx.Param("id")

	err := databases.DB.Table("addresses").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID, email, role")
	}).
		Where("id = ? AND user_id = ?", addressID, userId).First(&addresses).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Address not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve address",
			})

			return
		}
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Address detail retrieved successfully",
		"data":    addresses,
	})

}

func CreateAddress(ctx *gin.Context) {
	var input models.PayloadAddress
	var User models.Users
	user_id, _ := ctx.Get("userID")

	if errInput := ctx.ShouldBindJSON(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	regexNotelepon := regexp.MustCompile(`^\+62\s?\d{2,3}[-\s]?\d{3,4}[-\s]?\d{3,4}|\(0\d{2,3}\)\s?\d{3,4}[-\s]?\d{3,4}|\+62\d{8,14}|0\d{9,13}$`)
	if !regexNotelepon.MatchString(input.Phone) {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Invalid phone number format",
		})

		return
	}

	if err := databases.DB.Table("users").Where("id = ?", user_id).First(&User).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "User not found",
		})
		return
	}

	data := models.Address{
		UserID:       uint64(User.ID),
		Street:       input.Street,
		City:         input.City,
		State:        input.State,
		PostalCode:   input.PostalCode,
		Country:      input.Country,
		Phone:        input.Phone,
		NamaAlamat:   input.NamaAlamat,
		DetailAlamat: input.DetailAlamat,
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
	}

	if err := databases.DB.Table("addresses").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to create address",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "Create data Address success",
		"data":    input,
	})
}

func UpdateAddress(ctx *gin.Context) {
	var input models.PayloadAddress
	var address models.Address
	var User models.Users
	user_id, _ := ctx.Get("userID")
	addressId := ctx.Param("id")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	if err := databases.DB.Table("users").Where("id = ?", user_id).First(&User).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "User not found",
		})
		return
	}

	regexNotelepon := regexp.MustCompile(`^\+62\s?\d{2,3}[-\s]?\d{3,4}[-\s]?\d{3,4}|\(0\d{2,3}\)\s?\d{3,4}[-\s]?\d{3,4}|\+62\d{8,14}|0\d{9,13}$`)
	if !regexNotelepon.MatchString(input.Phone) {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Invalid phone number format",
		})

		return
	}

	// Cari address berdasarkan user_id dan addressId
	if err := databases.DB.Table("addresses").
		Where("id = ? AND user_id = ?", addressId, user_id).
		First(&address).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Address not found for this user",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve address",
			})

			return
		}
	}

	data := models.Address{
		UserID:       uint64(User.ID),
		Street:       input.Street,
		City:         input.City,
		State:        input.State,
		PostalCode:   input.PostalCode,
		Country:      input.Country,
		Phone:        input.Phone,
		NamaAlamat:   input.NamaAlamat,
		DetailAlamat: input.DetailAlamat,
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
	}

	if err := databases.DB.Table("addresses").Where("id = ? AND user_id = ?", addressId, user_id).Updates(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update address",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Address updated successfully",
		"data":    data,
	})
}

func DeleteAddress(ctx *gin.Context) {
	var address models.Address
	addressId := ctx.Param("id")
	userId, _ := ctx.Get("userID")

	if err := databases.DB.Table("addresses").Where("id = ? AND user_id = ?", addressId, userId).First(&address).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "Address not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve address",
			})

			return
		}
	}

	if err := databases.DB.Table("addresses").Where("id = ? AND user_id = ?", addressId, userId).Delete(&address).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to retrieve address",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Delete data Address success",
	})
}
