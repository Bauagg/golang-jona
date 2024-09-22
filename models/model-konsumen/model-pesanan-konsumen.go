package modelkonsumens

import (
	models "backend-jona-golang/models/model-global"

	"gorm.io/gorm"
)

type StatusPesanan string

const (
	Menunggu     StatusPesanan = "menunggu"
	Berhasil     StatusPesanan = "berhasil"
	PesananBatal StatusPesanan = "pesanan batal"
)

type PesananKonsumen struct {
	gorm.Model
	UserID              uint64             `json:"user_id" binding:"required"` // User ID (required)
	User                models.Users       `gorm:"foreignKey:UserID"`
	MetodePembayaran    uint64             `json:"metode_pembayaran" binding:"required"` // Payment Method ID (required)
	Bank                models.DaftarBank  `gorm:"foreignKey:MetodePembayaran"`
	JasaBersiId         uint64             `json:"jasa_bersi_id" binding:"required"` // Service ID (required)
	Jasa                models.SubCategory `gorm:"foreignKey:JasaBersiId"`
	CodePesanan         string             `json:"code_pesanan" binding:"required"`                                                                        // Order Code (required)
	Status              StatusPesanan      `json:"status" binding:"required" gorm:"type:enum('menunggu', 'berhasil', 'pesanan batal');default:'menunggu'"` // Order Status (required)
	TransactionMidtrans string             `json:"transaction_midtrans" binding:"required"`
	VaBank              string             `json:"va_bank" gorm:"null"`
}

type InputPesananKonsumen struct {
	gorm.Model
	MetodePembayaran uint64 `json:"metode_pembayaran" binding:"required"` // Payment Method ID (required)
	JasaBersiId      uint64 `json:"jasa_bersi_id" binding:"required"`     // Service ID (required)                                                                       // Order Code (required)
}
