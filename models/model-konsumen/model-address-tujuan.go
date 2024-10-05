package modelkonsumens

import "gorm.io/gorm"

type AddressTujuna struct {
	gorm.Model
	Street       string  `gorm:"size:255;not null" json:"street"`
	City         string  `gorm:"size:100;not null" json:"city"`
	State        string  `gorm:"size:100;not null" json:"state"`
	PostalCode   string  `gorm:"size:20;not null" json:"postal_code"`
	Country      string  `gorm:"size:100;not null" json:"country"`
	Phone        string  `gorm:"size:100;not null" json:"phone"`
	Nama         string  `gorm:"size:255;not null" json:"nama"`
	DetailAlamat string  `gorm:"size:255;not null" json:"detail_alamat"`
	Latitude     float64 `gorm:"type:decimal(9,6);not null" json:"latitude"`
	Longitude    float64 `gorm:"type:decimal(9,6);not null" json:"longitude"`
}

type PayloadAddressTujuan struct {
	Street       string  `gorm:"size:255;not null" json:"street"`
	City         string  `gorm:"size:100;not null" json:"city"`
	State        string  `gorm:"size:100;not null" json:"state"`
	PostalCode   string  `gorm:"size:20;not null" json:"postal_code"`
	Country      string  `gorm:"size:100;not null" json:"country"`
	Phone        string  `gorm:"size:100;not null" json:"phone"`
	Nama         string  `gorm:"size:255;not null" json:"nama"`
	DetailAlamat string  `gorm:"size:255;not null" json:"detail_alamat"`
	Latitude     float64 `gorm:"type:decimal(9,6);not null" json:"latitude"`
	Longitude    float64 `gorm:"type:decimal(9,6);not null" json:"longitude"`
}
