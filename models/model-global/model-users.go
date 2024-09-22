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
	Nama               string `json:"nama" binding:"required"`
	Email              string `json:"email" binding:"required" gorm:"unique"`                                                            // Wajib diisi
	Password           string `json:"password" binding:"required"`                                                                       // Wajib diisi, disembunyikan saat diubah ke JSON
	KonfirmasiPassword string `json:"konfirmasi_password" binding:"required" gorm:"-"`                                                   // Wajib diisi, tidak disimpan ke DB
	Role               Role   `json:"role" binding:"required" gorm:"type:enum('konsumen', 'jasa', 'driver', 'toko');default:'konsumen'"` // Wajib diisi
	Profile            string `json:"profile" gorm:"null"`
}

// Input Login User
type InputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Input Email User
type InputEmail struct {
	Email string `json:"email" binding:"required"`
}

// Input Password User
type InputPassword struct {
	Password           string `json:"password" binding:"required"` // Wajib diisi, disembunyikan saat diubah ke JSON
	KonfirmasiPassword string `json:"konfirmasi_password" binding:"required" gorm:"-"`
}
