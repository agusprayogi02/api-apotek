package routes

import (
	"api-apotek/controller"
	"api-apotek/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	user controller.UserController
}

func (r Routes) GetRoutes() *gin.Engine {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTAuth = service.JwtAuthService()
	var loginController = controller.LoginHandler(loginService, jwtService)

	router := gin.Default()

	v1 := router.Group("v1")
	v1.GET("/user", r.user.ViewAllUsers)
	v1.POST("/user/create", r.user.CreateUser)
	v1.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
		ctx.JSON(http.StatusUnauthorized, nil)
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":   "Welcome to gin gonic Backend",
			"status": http.StatusOK,
		})
	})
	return router
}
