package models

import "gorm.io/gorm"

type SubCategory struct {
	gorm.Model
	Nama  string `json:"nama" binding:"required"`
	Harga uint64 `json:"harga" binding:"required"`
}
