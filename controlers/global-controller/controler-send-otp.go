package controlers

import (
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"
	"backend-jona-golang/utils"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VerifyOTPPassword(ctx *gin.Context) {
	type InputOtp struct {
		NumberOtp uint64 `json:"number_otp" binding:"required" gorm:"unique"`
	}

	var input InputOtp

	// Bind input JSON to struct
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Check if OTP exists and is valid
	var otp models.OTP
	err := databases.DB.Table("otps").Where("number_otp = ? AND user_id = ? AND expires_at > ?", input.NumberOtp, ctx.Param("id"), time.Now()).First(&otp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid OTP or OTP has expired.",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Internal server error.",
			})

			return
		}
	}

	// If OTP is valid, return success
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "OTP verified successfully.",
	})
}

func VerifyOTP(ctx *gin.Context) {
	type InputOtp struct {
		NumberOtp uint64 `json:"number_otp" binding:"required" gorm:"unique"`
	}

	var input InputOtp
	userID, _ := ctx.Get("userID")

	// Bind input JSON to struct
	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})

		return
	}

	// Check if OTP exists and is valid
	var otp models.OTP
	err := databases.DB.Table("otps").Where("number_otp = ? AND user_id = ? AND expires_at > ?", input.NumberOtp, userID, time.Now()).First(&otp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid OTP or OTP has expired.",
			})

			return
		} else {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "Internal server error.",
			})

			return
		}
	}

	// If OTP is valid, return success
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "OTP verified successfully.",
	})
}

func SendEmailOtp(ctx *gin.Context) {
	email, _ := ctx.Get("userEmail")
	userId, _ := ctx.Get("userID")

	// Generate a random 4-digit OTP
	rand.Seed(time.Now().UnixNano())
	randomOTP := rand.Intn(9000) + 1000

	errSendEmail := utils.SendEmail(email.(string), uint64(randomOTP))
	if errSendEmail != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to send OTP email.",
		})

		return
	}

	// Check if OTP already exists
	var existingOtp models.OTP
	errFindOtp := databases.DB.Table("otps").Where("user_id = ?", userId).First(&existingOtp).Error
	if errFindOtp != nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "OTP not found for the user.",
		})
		return
	}

	// OTP exists, update it
	existingOtp.NumberOtp = uint64(randomOTP)
	existingOtp.ExpiresAt = time.Now().Add(5 * time.Minute)

	errUpdateOtp := databases.DB.Table("otps").Where("user_id = ? AND id = ?", existingOtp.UserId, existingOtp.ID).Updates(models.OTP{
		NumberOtp: uint64(randomOTP),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}).Error

	if errUpdateOtp != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to update OTP.",
			"data":    existingOtp,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "OTP updated successfully.",
	})
}
