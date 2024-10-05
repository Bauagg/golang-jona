package migrate

import (
	"backend-jona-golang/databases"
	models "backend-jona-golang/models/model-global"
	modelkonsumens "backend-jona-golang/models/model-konsumen"
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
		&models.DaftarBank{},
		&models.CategorySpecific{},

		// Tambahkan model konsumen
		&modelkonsumens.PesananKonsumen{},
		&modelkonsumens.NotifikasiPembayaran{},
		&modelkonsumens.AddressTujuna{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
