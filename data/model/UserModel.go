package model

import (
	"api-apotek/config"
	"api-apotek/data/entity"
	"fmt"
)

type UserModel struct {
	db config.DatabaseConfig
}

func (e *UserModel) GetUsers() (*[]entity.User, error) {
	var users []entity.User
	err := e.db.GetDatabase().Find(&users).Error
	if err != nil {
		fmt.Printf("[userModel.getUsers] error :%v", err)
		return nil, fmt.Errorf("failed get all users")
	}
	return &users, nil
}

func (e *UserModel) AddUsers(user *entity.User) (*entity.User, error) {
	err := e.db.GetDatabase().Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return user, nil
}
