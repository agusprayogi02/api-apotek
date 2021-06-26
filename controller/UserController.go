package controller

import (
	"api-apotek/data/entity"
	"api-apotek/data/model"
	"api-apotek/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userModel model.UserModel
}

func (e *UserController) ViewAllUsers(c *gin.Context) {
	users, err := e.userModel.GetUsers()
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "Not Found")
		return
	}
	utils.Success(c, users)
}

func (e *UserController) CreateUser(c *gin.Context) {
	var user = entity.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.Nama == "" || user.Email == "" || user.Password == "" {
		utils.HandleError(c, http.StatusBadRequest, "Field are Required!")
		return
	}
	user.Password = utils.Encrypt(user.Password, utils.GetRandKey())
	mUser, err := e.userModel.AddUsers(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, mUser)
}
