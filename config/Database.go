package config

import (
	"api-apotek/entity"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type DatabaseConfig struct {
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "",
		DBName:   "apotek_db",
	}
	return &dbConfig
}

func (DatabaseConfig DatabaseConfig) GetDatabase() *gorm.DB {
	dbConfig := BuildDBConfig()
	config := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	DB, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func MigrateDB(db *gorm.DB) {
	db.Debug().AutoMigrate(
		entity.User{},
		entity.Toko{},
	)
}
