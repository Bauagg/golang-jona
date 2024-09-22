package models

import "gorm.io/gorm"

type StatusBank string

const (
	TranferBank StatusBank = "Transfer Bank"
	Tunai       StatusBank = "Tunai"
)

type DaftarBank struct {
	gorm.Model
	Nama   string     `json:"nama" form:"nama" binding:"required"`
	Icon   string     `json:"icon" form:"nama"`
	Status StatusBank `json:"role" form:"status" binding:"required" gorm:"type:enum('Transfer Bank', 'Tunai');"` // Wajib diisi
}
