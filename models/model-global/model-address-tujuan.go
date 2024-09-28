package models

import "gorm.io/gorm"

type AddressTujuna struct {
	gorm.Model
	UserID       uint64  `json:"user_id" binding:"required"` // Foreign key linking to a User, if applicable
	Street       string  `gorm:"size:255;not null" json:"street"`
	City         string  `gorm:"size:100;not null" json:"city"`
	State        string  `gorm:"size:100;not null" json:"state"`
	PostalCode   string  `gorm:"size:20;not null" json:"postal_code"`
	Country      string  `gorm:"size:100;not null" json:"country"`
	Phone        string  `gorm:"size:100;not null" json:"phone"`
	NamaAlamat   string  `gorm:"size:255;not null" json:"nama_alamat"`
	DetailAlamat string  `gorm:"size:255;not null" json:"detail_alamat"`
	Latitude     float64 `gorm:"type:decimal(9,6);not null" json:"latitude"`  // Latitude with 6 decimal places
	Longitude    float64 `gorm:"type:decimal(9,6);not null" json:"longitude"` // Longitude with 6 decimal places
	Status       string  `gorm:"size:255;null" json:"status"`
}

type PayloadAddressTujuan struct {
	Street       string  `gorm:"size:255;not null" json:"street"`
	City         string  `gorm:"size:100;not null" json:"city"`
	State        string  `gorm:"size:100;not null" json:"state"`
	PostalCode   string  `gorm:"size:20;not null" json:"postal_code"`
	Country      string  `gorm:"size:100;not null" json:"country"`
	Phone        string  `gorm:"size:100;not null" json:"phone"`
	NamaAlamat   string  `gorm:"size:255;not null" json:"nama_alamat"`
	DetailAlamat string  `gorm:"size:255;not null" json:"detail_alamat"`
	Latitude     float64 `gorm:"type:decimal(9,6);not null" json:"latitude"`  // Latitude with 6 decimal places
	Longitude    float64 `gorm:"type:decimal(9,6);not null" json:"longitude"` // Longitude with 6 decimal places
}
