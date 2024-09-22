package konsumencontrollers

import (
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"

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
