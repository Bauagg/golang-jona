package controlers

import (
	"backend-jona-golang/config"
	"backend-jona-golang/databases"
	"backend-jona-golang/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFiturJona(ctx *gin.Context) {
	var fiturJona []models.FiturJona

	if err := databases.DB.Table("fitur_jonas").Find(&fiturJona).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to retrieve icons",
		})

		return
	}

	// Return success response
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data fitur jona success",
		"data":    fiturJona,
	})
}

func CreateFitur(ctx *gin.Context) {
	var input models.FiturJona

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Get the uploaded file from form data
	file, err := ctx.FormFile("icon")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to upload image: " + err.Error(),
		})
		return
	}

	// Create directory if not exists
	imageDir := "./public/image-fitur"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		err = os.MkdirAll(imageDir, os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": "Failed to create image directory: " + err.Error(),
			})
			return
		}
	}

	// Generate unique filename based on timestamp
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filepath := filepath.Join(imageDir, fileName)

	// Save the file to the server
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to save image: " + err.Error(),
		})
		return
	}

	// Set the icon URL (relative path to the file)
	data := models.FiturJona{
		Icon: config.URL_HOST + "/images/" + fileName,
		Nama: input.Nama,
	}

	// Save the feature to the database (assuming you have a DB instance setup)
	if err := databases.DB.Table("fitur_jonas").Create(&data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to save feature to database: " + err.Error(),
		})
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Feature created successfully",
		"data":    input,
	})
}

func UpdateFiturJona(ctx *gin.Context) {
	var input models.FiturJona
	var data models.FiturJona
	id := ctx.Param("id")

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Get the uploaded file from form data
	file, err := ctx.FormFile("icon")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to upload image: " + err.Error(),
		})
		return
	}

	if err := databases.DB.Table("fitur_jonas").Where("id = ?", id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data fitur_jonas not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve fitur_jonas",
			})

			return
		}
	}

	// Hapus file lama jika ada
	if data.Icon != "" {
		// Ambil nama file dari URL
		fileName := filepath.Base(data.Icon)
		oldFilePath := filepath.Join("./public/image-fitur", fileName)

		// Hapus file lama
		if err := os.Remove(oldFilePath); err != nil && !os.IsNotExist(err) {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": "Failed to delete old image: " + err.Error(),
			})
			return
		}
	}

	// Create directory if not exists
	imageDir := "./public/image-fitur"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		err = os.MkdirAll(imageDir, os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": "Failed to create image directory: " + err.Error(),
			})
			return
		}
	}

	// Generate unique filename based on timestamp
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filepath := filepath.Join(imageDir, fileName)

	// Save the file to the server
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to save image: " + err.Error(),
		})
		return
	}

	data.Nama = input.Nama
	data.Icon = config.URL_HOST + "/images/" + fileName

	// update data fitur jona
	if err := databases.DB.Table("fitur_jonas").Where("id = ?", id).Updates(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update fitur_jonas",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "Update data success",
		"data":    data,
	})
}

func DeleteFiturJona(ctx *gin.Context) {
	var data models.FiturJona
	id := ctx.Param("id")

	if err := databases.DB.Table("fitur_jonas").Where("id = ?", id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data fitur_jonas not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve fitur_jonas",
			})

			return
		}
	}

	if err := databases.DB.Table("fitur_jonas").Where("id = ?", id).Delete(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to delete fitur_jonas",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Delete data Fitur_jona success",
	})
}
