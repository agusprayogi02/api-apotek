package config

import (
	"log"

	"api-apotek/entity"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type DatabaseConfig struct{}

func (DatabaseConfig DatabaseConfig) getDatabaseConfig() *gorm.DB {
	DB, err := gorm.Open("mysql", viper.GetString("database.mysql"))
	if err != nil {
		log.Fatal(err)
	}
	DB.Debug().AutoMigrate(
		entity.User{},
	)
	return DB
}
