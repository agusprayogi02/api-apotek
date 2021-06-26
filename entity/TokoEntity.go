package entity

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	Nama       string `gorm:"not null;size:150" json:"nama"`
	Lokasi     string `gorm:"not null;size:225" json:"lokasi"`
	Coordinate string `gorm:"not null;size:100" json:"coordinate"` // coordinate
}

func (e Toko) TableName() string {
	return "toko"
}
