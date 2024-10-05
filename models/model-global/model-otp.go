package models

import (
	"time"

	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model
	NumberOtp uint64    `json:"number_otp" form:"number_otp" binding:"required" gorm:"unique"`
	UserId    uint64    `json:"user_id" form:"user_id" binding:"required"`
	ExpiresAt time.Time `json:"expires_at" form:"expires_at" binding:"required"` // Field to store expiration time
}
