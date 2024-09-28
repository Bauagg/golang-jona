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
	OrderID       string `json:"order_id" binding:"required"`
	UserId        uint64 `json:"user_id" binding:"required"`
}

// MidtransNotification represents the structure of the Midtrans notification
type MidtransNotification struct {
	TransactionTime        string `json:"transaction_time"`
	TransactionStatus      string `json:"transaction_status"`
	TransactionID          string `json:"transaction_id"`
	StatusMessage          string `json:"status_message"`
	StatusCode             string `json:"status_code"`
	SignatureKey           string `json:"signature_key"`
	PaymentType            string `json:"payment_type"`
	OrderID                string `json:"order_id"`
	MerchantID             string `json:"merchant_id"`
	MaskedCard             string `json:"masked_card"`
	GrossAmount            string `json:"gross_amount"`
	FraudStatus            string `json:"fraud_status"`
	Eci                    string `json:"eci"`
	Currency               string `json:"currency"`
	ChannelResponseMessage string `json:"channel_response_message"`
	ChannelResponseCode    string `json:"channel_response_code"`
	CardType               string `json:"card_type"`
	Bank                   string `json:"bank"`
	ApprovalCode           string `json:"approval_code"`
}
