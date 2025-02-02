package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	Konsumen Role = "konsumen"
	Jasa     Role = "jasa"
	Driver   Role = "driver"
	Toko     Role = "toko"
)

// model databases
type Users struct {
	gorm.Model
	Email              string `json:"email" binding:"required" gorm:"unique"`                                                            // Wajib diisi
	Password           string `json:"password" binding:"required"`                                                                       // Wajib diisi, disembunyikan saat diubah ke JSON
	KonfirmasiPassword string `json:"konfirmasi_password" binding:"required" gorm:"-"`                                                   // Wajib diisi, tidak disimpan ke DB
	Role               Role   `json:"role" binding:"required" gorm:"type:enum('konsumen', 'jasa', 'driver', 'toko');default:'konsumen'"` // Wajib diisi
}

// Input Login User
type InputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
