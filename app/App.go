package app

import (
	"api-apotek/config"
	"api-apotek/routes"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	config.DatabaseConfig.GetDatabase(config.DatabaseConfig{})

	app := routes.Routes.GetRoutes(routes.Routes{})
	return app
}
