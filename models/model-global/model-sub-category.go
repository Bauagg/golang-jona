package models

import "gorm.io/gorm"

type SubCategory struct {
	gorm.Model
	Nama         string       `json:"nama" binding:"required"`
	Harga        *uint64      `json:"harga" gorm:"default:null"`
	Description  *string      `json:"description" gorm:"default:null"`
	IdCategory   uint64       `json:"id_category" binding:"required"`
	CaegoryUtama CaegoryUtama `gorm:"foreignKey:IdCategory"`
}

type InputSubCategory struct {
	Nama        string `json:"nama" binding:"required"`
	Harga       uint64 `json:"harga"`
	Description string `json:"description"`
	IdCategory  uint64 `json:"id_category" binding:"required"`
}
