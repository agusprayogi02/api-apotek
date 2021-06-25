package model

import (
	"api-apotek/config"
	"api-apotek/entity"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	db config.DatabaseConfig
}

func (e *UserModel) GetUsers(c *gin.Context) (*[]entity.User, error) {
	var users []entity.User
	err := e.db.GetDatabase().Find(&users).Error
	if err != nil {
		fmt.Printf("[userModel.getUsers] error :%v", err)
		return nil, fmt.Errorf("failed get all users")
	}
	return &users, nil
}
