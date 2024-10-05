package konsumencontrollers

import (
	"backend-jona-golang/config"
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileKonsumen(ctx *gin.Context) {
	var dataUser models.Users
	userId, _ := ctx.Get("userID")

	if err := databases.DB.Table("users").Where("id = ?", userId).First(&dataUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "User not found",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Failed to retrieve user profile",
			})

			return
		}
	}

	// Berikan respons JSON jika user ditemukan
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "User profile retrieved successfully",
		"data":    dataUser, // Kirimkan data user dalam respons
	})
}

func UpdateProfileKonsumen(ctx *gin.Context) {
	var input models.InputUpdateProfilekonsumen
	var data models.Users

	// Bind input data
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	// Get user ID from context
	user_id, _ := ctx.Get("userID")

	// Fetch user data based on user ID
	if err := databases.DB.Table("users").Where("id = ?", user_id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "User not found",
			})
			return
		}
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Database User error",
		})
		return
	}

	// Update fields if provided in input
	if input.Email != "" {
		// Validasi format email
		regexEmaill := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !regexEmaill.MatchString(input.Email) {
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid email format",
			})

			return
		}

		data.Email = input.Email

		err := databases.DB.Table("users").Where("email = ?", input.Email).First(&data).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Internal server error",
			})
			return
		}

		if err == nil {
			// Jika tidak ada error dan user ditemukan, email sudah terdaftar
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "Email sudah terdaftar.",
			})
			return
		}
	}

	if input.Nama != "" {
		data.Nama = input.Nama
	}

	// Handle profile picture update if a file is uploaded
	file, _ := ctx.FormFile("profile")
	if file != nil {
		imageDir := "./public/profile-user"

		// Delete old profile image if exists
		if data.Profile != "" {
			fileName := filepath.Base(data.Profile)
			oldFilePath := filepath.Join(imageDir, fileName)

			if err := os.Remove(oldFilePath); err != nil && !os.IsNotExist(err) {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   true,
					"message": "Failed to delete old image: " + err.Error(),
				})
				return
			}
		}

		// Create directory if not exists
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

		// Save new profile image
		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := filepath.Join(imageDir, fileName)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": "Failed to save image: " + err.Error(),
			})
			return
		}

		// Update profile field with new image URL
		data.Profile = config.URL_HOST + "/profile-user/" + fileName
	}

	// Save the updated user data to the database
	if err := databases.DB.Table("users").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update user profile",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Update Profile User success",
		"data":    data,
	})
}
