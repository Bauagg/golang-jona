package konsumencontrollers

import (
	"backend-jona-golang/databases"
	modelkonsumens "backend-jona-golang/models/model-konsumen"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DetailAddressTujuan(ctx *gin.Context) {
	var data modelkonsumens.AddressTujuna

	if err := databases.DB.Table("address_tujunas").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data id Address Tujian Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databaese Address Tujian Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "detail Address tujuan success",
		"data":    data,
	})
}

func CreateAddressTujuan(ctx *gin.Context) {
	var input modelkonsumens.PayloadAddressTujuan
	var data modelkonsumens.AddressTujuna

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	data.Street = input.Street
	data.City = input.City
	data.State = input.State
	data.PostalCode = input.PostalCode
	data.Country = input.Country
	data.Phone = input.Phone
	data.Nama = input.Nama
	data.DetailAlamat = input.DetailAlamat
	data.Latitude = input.Latitude
	data.Longitude = input.Longitude

	if err := databases.DB.Table("address_tujunas").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databases Address Tujuan error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "create data Alamat Tujuan success",
		"data":    data,
	})
}

func UpdateAddressTujuan(ctx *gin.Context) {
	var input modelkonsumens.PayloadAddressTujuan
	var data modelkonsumens.AddressTujuna

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("address_tujunas").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data id Address Tujian Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databaese Address Tujian Error",
		})
		return
	}

	data.Street = input.Street
	data.City = input.City
	data.State = input.State
	data.PostalCode = input.PostalCode
	data.Country = input.Country
	data.Phone = input.Phone
	data.Nama = input.Nama
	data.DetailAlamat = input.DetailAlamat
	data.Latitude = input.Latitude
	data.Longitude = input.Longitude

	if err := databases.DB.Table("address_tujunas").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "databases Address Tujuan error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "create data Alamat Tujuan success",
		"data":    data,
	})
}
