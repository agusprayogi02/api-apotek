package controller

import (
	"api-apotek/model"
	"api-apotek/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userModel model.UserModel
}

func (e *UserController) ViewAllUsers(c *gin.Context) {
	users, err := e.userModel.GetUsers(c)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "Not Found")
	}
	utils.Success(c, users)
}
