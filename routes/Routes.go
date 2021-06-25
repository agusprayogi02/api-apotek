package routes

import (
	"api-apotek/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	user controller.UserController
}

func (r Routes) GetRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("v1")
	v1.GET("/user", r.user.ViewAllUsers)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":   "Welcome to gin gonic Backend",
			"status": http.StatusOK,
		})
	})
	return router
}
