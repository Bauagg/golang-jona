package modelkonsumens

import (
	models "backend-jona-golang/models/model-global"

	"gorm.io/gorm"
)

type StatusPesanan string

const (
	Menunggu     StatusPesanan = "menunggu"
	Berhasil     StatusPesanan = "berhasil"
	PesananBatal StatusPesanan = "dibatalkan"
	Kadaluarsa   StatusPesanan = "kadaluwarsa"
	ErrorPesanan StatusPesanan = "error dalam pesanan"
)

type PesananKonsumen struct {
	gorm.Model
	UserID              uint64             `json:"user_id" binding:"required"` // User ID (required)
	User                models.Users       `gorm:"foreignKey:UserID"`
	MetodePembayaran    uint64             `json:"metode_pembayaran" binding:"required"` // Payment Method ID (required)
	Bank                models.DaftarBank  `gorm:"foreignKey:MetodePembayaran"`
	JasaId              uint64             `json:"jasa_id" binding:"required"` // Service ID (required)
	Jasa                models.SubCategory `gorm:"foreignKey:JasaId"`
	CodePesanan         string             `json:"code_pesanan" binding:"required"`                                                                                                           // Order Code (required)
	Status              StatusPesanan      `json:"status" binding:"required" gorm:"type:enum('menunggu', 'berhasil', 'dibatalkan', 'error dalam pesanan', 'kadaluwarsa');default:'menunggu'"` // Order Status (required)
	TransactionMidtrans string             `json:"transaction_midtrans" binding:"required"`
	IdAddress           uint64             `json:"id_address" binding:"required"`
	Address             models.Address     `gorm:"foreignKey:IdAddress"`
	VaBank              string             `json:"va_bank" gorm:"null"`
	AlamatTujuan        uint64             `json:"alamat_tujuan" gorm:"null"`
	SpecificCategory    uint64             `json:"specific_category" gorm:"null"`
}

type InputPesananKonsumen struct {
	MetodePembayaran uint64 `json:"metode_pembayaran" binding:"required"` // Payment Method ID (required)
	JasaId           uint64 `json:"jasa_id" binding:"required"`           // Service ID (required)
	IdAddress        uint64 `json:"id_address" binding:"required"`        // Order Code (required)
}

type InputPesananJasaKirimKonsumen struct {
	MetodePembayaran uint64 `json:"metode_pembayaran" binding:"required"` // Payment Method ID (required)
	JasaId           uint64 `json:"jasa_id" binding:"required"`           // Service ID (required)
	IdAddress        uint64 `json:"id_address" binding:"required"`        // Order Code (required)
	IdAlamatTujuan   uint64 `json:"id_alamat_tujuan" binding:"required"`
	SpecificCategory uint64 `json:"specific_category" binding:"required"`
}
