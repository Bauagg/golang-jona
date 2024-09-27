package modelkonsumens

import "gorm.io/gorm"

type Status string

const (
	NotifikasiMenunggu Status = "Menunggu Konfirmasi Pembayaran"
	NotifikasiBatal    Status = "Pesanan Batal"
	NotifikasiBerhasil Status = "Pembayaran Berhasil"
)

type NotifikasiPembayaran struct {
	gorm.Model
	StatusPesanan Status `json:"status_pesanan" binding:"required" gorm:"type:enum('Menunggu Konfirmasi Pembayaran', 'Pesanan Batal', 'Pembayaran Berhasil');default:'Menunggu Konfirmasi Pembayaran'"`
	Description   string `json:"description" binding:"required"`
	TransactionID string `json:"transaction_id" binding:"required"`
	UserId        uint64 `json:"user_id" binding:"required"`
}
