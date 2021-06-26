package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(100) not null"`
	Email    string `gorm:"uniqueIndex:idx_email;type:varchar(100) not null"`
	Password string `gorm:"size:225;not null"`
}

func (e User) TableName() string {
	return "users"
}
