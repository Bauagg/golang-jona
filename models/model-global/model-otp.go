package models

import (
	"time"

	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model
	NumberOtp uint64    `json:"number_otp" form:"nama" binding:"required" gorm:"unique"`
	UserId    uint64    `json:"user_id" form:"nama" binding:"required"`
	ExpiresAt time.Time `json:"expires_at" form:"nama" binding:"required"` // Field to store expiration time
}
