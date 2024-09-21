package migrate

import (
	"backend-jona-golang/databases"
	"backend-jona-golang/models"
)

func Migrate() {
	db := databases.GetDB()
	err := db.AutoMigrate(
		// Tambahkan model baru di sini
		&models.Users{},
		&models.OTP{},
		&models.Address{},
		&models.FiturJona{},
		&models.CaegoryUtama{},
		&models.SubCategory{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
