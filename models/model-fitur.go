package models

import "gorm.io/gorm"

type FiturJona struct {
	gorm.Model
	Nama string `json:"nama" form:"nama" binding:"required"`
	Icon string `json:"icon" form:"nama" binding:"required"`
}
