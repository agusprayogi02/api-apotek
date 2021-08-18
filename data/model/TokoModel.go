package model

import (
	"api-apotek/config"
	"api-apotek/data/entity"
)

type TokoModel struct {
	db config.DatabaseConfig
}

func (e *TokoModel) GetToko() (*[]entity.Toko, error) {
	var shops []entity.Toko
	err := e.db.GetDatabase().Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return &shops, nil
}
