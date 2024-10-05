package models

import "gorm.io/gorm"

type CategorySpecific struct {
	gorm.Model
	Price         uint64 `json:"price" binding:"required"`           // Wajib diisi
	Nama          string `json:"nama" binding:"required"`            // Wajib diisi
	Description   string `json:"description" binding:"required"`     // Wajib diisi
	IdSubCategory uint64 `json:"id_sub_category" binding:"required"` // Wajib diisi
}
