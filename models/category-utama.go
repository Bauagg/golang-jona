package models

import "gorm.io/gorm"

type CaegoryUtama struct {
	gorm.Model
	FiturId     uint64    `json:"fitur_id" binding:"required"`
	Fitur       FiturJona `gorm:"foreignKey:FiturId"`
	Nama        string    `json:"nama" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

type InputCaegoryUtama struct {
	gorm.Model
	FiturId     uint64 `json:"fitur_id" binding:"required"`
	Nama        string `json:"nama" binding:"required"`
	Description string `json:"description" binding:"required"`
}
