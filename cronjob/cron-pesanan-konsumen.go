package cronjob

import (
	"backend-jona-golang/databases"
	modelkonsumens "backend-jona-golang/models/model-konsumen"
	"time"
)

func UpdateExpiredOrders() {
	var pesanan []modelkonsumens.PesananKonsumen

	// Cari pesanan yang statusnya "menunggu" dan dibuat lebih dari 30 menit yang lalu
	expirationTime := time.Now().Add(-30 * time.Minute)

	databases.DB.Table("pesanan_konsumens").
		Where("status = ? AND created_at < ?", modelkonsumens.Menunggu, expirationTime).
		Find(&pesanan)

		// Update status menjadi "pesanan batal" untuk pesanan yang kedaluwarsa
	for _, order := range pesanan {
		order.Status = modelkonsumens.Kadaluarsa
		databases.DB.Save(&order)

		newNotification := modelkonsumens.NotifikasiPembayaran{
			Description:   "Waktu Pembayaran kamu habis",
			UserId:        order.UserID,
			StatusPesanan: modelkonsumens.NotifikasiBatal,
			TransactionID: order.TransactionMidtrans,
		}

		databases.DB.Table("notifikasi_pembayarans").Where("transaction_id = ?", order.TransactionMidtrans).Updates(&newNotification)
	}
}
