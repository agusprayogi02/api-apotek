package controller

import (
	"api-apotek/dto"
	"api-apotek/service"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Login(ctx *gin.Context) string
}

type LoginController struct {
	LoginService service.LoginService
	JWTAuth      service.JWTAuth
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTAuth) Controller {
	return &LoginController{
		LoginService: loginService,
		JWTAuth:      jwtService,
	}
}

func (controller *LoginController) Login(ctx *gin.Context) string {
	var credential *dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return err.Error()
	}
	isUserAuthenticated := controller.LoginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.JWTAuth.GenerateToken(credential.Email, true)
	}
	return ""
}
