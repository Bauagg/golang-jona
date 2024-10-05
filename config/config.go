package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST               string
	DB_PORT               string
	DB_NAME               string
	DB_USER               string
	DB_PASSWORD           string
	APP_PORT              string
	SECRET_KEY            string
	GMAIL_OTP             string
	PASSWORD_OTP          string
	URL_HOST              string
	URL_MIDTRANS          string
	SERVER_KEY_MIDTRANS   string
	URL_NOTIFIKASI        string
	APP_ID                string
	SERVER_KEY_NOTIFIKASI string
)

func InitConfigEnv() {
	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, proceeding with environment variables")
	}

	// Fetch values from environment variables
	appPortEnv := os.Getenv("APP_PORT")
	if appPortEnv != "" {
		APP_PORT = appPortEnv
	}

	dbHostEnv := os.Getenv("DB_HOST")
	if dbHostEnv != "" {
		DB_HOST = dbHostEnv
	}

	dbNameEnv := os.Getenv("DB_NAME")
	if dbNameEnv != "" {
		DB_NAME = dbNameEnv
	}

	dbPasswordEnv := os.Getenv("DB_PASSWORD")
	if dbPasswordEnv != "" {
		DB_PASSWORD = dbPasswordEnv
	}

	dbPortEnv := os.Getenv("DB_PORT")
	if dbPortEnv != "" {
		DB_PORT = dbPortEnv
	}

	dbUserEnv := os.Getenv("DB_USER")
	if dbUserEnv != "" {
		DB_USER = dbUserEnv
	}

	secretKeyEnv := os.Getenv("SECRET_KEY")
	if secretKeyEnv != "" {
		SECRET_KEY = secretKeyEnv
	}

	gmailOtp := os.Getenv("GMAIL_OTP")
	if gmailOtp != "" {
		GMAIL_OTP = gmailOtp
	}

	passwordPtp := os.Getenv("PASSWORD_OTP")
	if passwordPtp != "" {
		PASSWORD_OTP = passwordPtp
	}

	urlhost := os.Getenv("URL_HOST")
	if urlhost != "" {
		URL_HOST = urlhost
	}

	// Midtrans
	urlMidtrans := os.Getenv("URL_MIDTRANS")
	if urlMidtrans != "" {
		URL_MIDTRANS = urlMidtrans
	}

	serverKeyMidtrans := os.Getenv("SERVER_KEY_MIDTRANS")
	if serverKeyMidtrans != "" {
		SERVER_KEY_MIDTRANS = serverKeyMidtrans
	}

	// Notifikasi
	urlNotifikasi := os.Getenv("URL_NOTIFIKASI")
	if urlNotifikasi != "" {
		URL_NOTIFIKASI = urlNotifikasi
	}

	appID := os.Getenv("APP_ID")
	if appID != "" {
		APP_ID = appID
	}

	serverKeyNotifikasi := os.Getenv("SERVER_KEY_NOTIFIKASI")
	if serverKeyNotifikasi != "" {
		SERVER_KEY_NOTIFIKASI = serverKeyNotifikasi
	}
}
