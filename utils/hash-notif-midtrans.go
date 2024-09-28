package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

// Fungsi untuk menghitung hash SHA512
func GenerateMidtransSignature(orderID, statusCode, grossAmount, serverKey string) string {
	signatureString := orderID + statusCode + grossAmount + serverKey

	// Menghasilkan hash SHA512
	hash := sha512.New()
	hash.Write([]byte(signatureString))
	hashedBytes := hash.Sum(nil)

	// Encode hasil hash menjadi string hexadecimal
	signature := hex.EncodeToString(hashedBytes)
	return signature
}
