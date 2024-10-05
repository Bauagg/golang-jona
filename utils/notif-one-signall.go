package utils

import (
	"backend-jona-golang/config"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// Struct untuk body request ke OneSignal
type NotificationRequestBody struct {
	AppID            string            `json:"app_id"`
	Contents         map[string]string `json:"contents"`
	IncludedSegments []string          `json:"include_external_user_ids"`
	Headings         map[string]string ` json:"headings"`
}

// Fungsi untuk mengirim notifikasi pembayaran menggunakan Resty
func SendPaymentNotification(message string, heading string, userids []string) error {
	client := resty.New()

	// Data notifikasi yang akan dikirim
	requestBody := NotificationRequestBody{
		AppID:            config.APP_ID,                    // Ganti dengan App ID dari OneSignal
		Contents:         map[string]string{"en": message}, // Isi notifikasi
		IncludedSegments: userids,
		Headings:         map[string]string{"en": heading}, // User ID yang akan menerima notifikasi
	}

	var serverKeyNotifikasi string = "Basic " + config.SERVER_KEY_NOTIFIKASI

	// Kirim request POST ke OneSignal
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", serverKeyNotifikasi). // Ganti dengan REST API Key dari OneSignal
		SetBody(requestBody).
		Post(config.URL_NOTIFIKASI)

	if err != nil {
		log.Print(err)
		return err
	}

	// Cek status code apakah berhasil atau tidak
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		log.Printf("Response from OneSignal: %s", resp.String())
		return fmt.Errorf("gagal mengirim notifikasi, status code: %d", resp.StatusCode())
	}

	return nil
}
