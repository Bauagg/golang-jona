package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOrderID menghasilkan OrderID dengan format JONA123456789
func GenerateOrderID(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(10000000000)             // Angka acak antara 0 dan 999999999
	return fmt.Sprintf("%s%09d", prefix, randomNum) // Menambahkan angka dengan format 9 digit
}
