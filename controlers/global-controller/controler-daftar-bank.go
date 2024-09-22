package controlers

import (
	"backend-jona-golang/config"
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListDaftarBankAll(ctx *gin.Context) {
	var data []models.DaftarBank

	if err := databases.DB.Table("daftar_banks").Find(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to retrieve daftar Bank",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data fitur jona success",
		"data":    data,
	})
}

func CreateDaftarBank(ctx *gin.Context) {
	var input models.DaftarBank

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	file, err := ctx.FormFile("icon")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to upload image: " + err.Error(),
		})

		return
	}

	imageDir := "./public/image-bank"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		err = os.MkdirAll(imageDir, os.ModePerm)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to create image directory: " + err.Error(),
			})
			return
		}
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filepath := filepath.Join(imageDir, fileName)

	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to save image: " + err.Error(),
		})
		return
	}

	input.Icon = config.URL_HOST + "/images-bank/" + fileName

	if err := databases.DB.Table("daftar_banks").Create(&input).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to save feature to database: " + err.Error(),
		})
		return
	}

	// Return success response
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Feature created successfully",
		"data":    input,
	})
}

func UpdateDaftarBank(ctx *gin.Context) {
	var input models.DaftarBank
	var data models.DaftarBank

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	file, err := ctx.FormFile("icon")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to upload image: " + err.Error(),
		})

		return
	}

	if err := databases.DB.Table("daftar_banks").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "No subcategories found for the given Daftar Bank ID",
			})
			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "An error occurred while retrieving the Daftar Bank: " + err.Error(),
			})

			return
		}
	}

	imageDir := "./public/image-bank"

	if data.Icon != "" {
		fileName := filepath.Base(data.Icon)
		oldFilePath := filepath.Join(imageDir, fileName)

		if err := os.Remove(oldFilePath); err != nil && !os.IsNotExist(err) {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to delete old image: " + err.Error(),
			})
			return
		}
	}

	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		err = os.MkdirAll(imageDir, os.ModePerm)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to create image directory: " + err.Error(),
			})
			return
		}
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filepath := filepath.Join(imageDir, fileName)

	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to save image: " + err.Error(),
		})
		return
	}

	data.Nama = input.Nama
	data.Icon = config.URL_HOST + "/images-bank/" + fileName
	data.Status = input.Status

	if err := databases.DB.Table("daftar_banks").Where("id = ?", ctx.Param("id")).Updates(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update Bank",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Update data success",
		"data":    data,
	})
}

func DeleteDaftarBank(ctx *gin.Context) {
	var data models.DaftarBank

	if err := databases.DB.Table("daftar_banks").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "No subcategories found for the given Daftar Bank ID",
			})
			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "An error occurred while retrieving the Daftar Bank: " + err.Error(),
			})

			return
		}
	}

	if data.Icon != "" {
		fileName := filepath.Base(data.Icon)
		oldFilePath := filepath.Join("./public/image-bank", fileName)

		if err := os.Remove(oldFilePath); err != nil && !os.IsNotExist(err) {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to delete old image: " + err.Error(),
			})
			return
		}
	}

	if err := databases.DB.Table("daftar_banks").Where("id = ?", ctx.Param("id")).Delete(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to delete Bank",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Delete data success",
	})
}
