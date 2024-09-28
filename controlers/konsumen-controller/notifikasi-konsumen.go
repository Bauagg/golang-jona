package konsumencontrollers

import (
	"backend-jona-golang/databases"
	modelkonsumens "backend-jona-golang/models/model-konsumen"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

var clients = make(map[string]chan string)

type NotificationDataTransaksi struct {
	OrderID       string `json:"order_id"`
	TransactionID string `json:"transaction_id"`
}

func SSEHandler(ctx *gin.Context) {
	channel := make(chan string)

	// Simpan channel untuk klien
	clients[ctx.Request.RemoteAddr] = channel

	// Set header untuk SSE
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	for {
		select {
		case message := <-channel:
			// Parse message to extract order and transaction IDs
			var notificationData NotificationDataTransaksi
			err := json.Unmarshal([]byte(message), &notificationData)
			if err != nil {
				ctx.SSEvent("error", "Failed to parse notification data")
				continue
			}

			// Mengambil data pesanan
			var notification modelkonsumens.NotifikasiPembayaran
			if err := databases.DB.Table("notifikasi_pembayarans").
				Where("transaction_id = ? AND order_id = ?", notificationData.TransactionID, notificationData.OrderID).
				First(&notification).Error; err != nil {
				ctx.SSEvent("error", "Order not found")
				continue
			}

			// Kirim data ke klien
			ctx.SSEvent("message", notification)

		case <-time.After(30 * time.Second):
			// Tutup koneksi setelah 30 detik tidak ada aktivitas
			return
		}
	}
}
