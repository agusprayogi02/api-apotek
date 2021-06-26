package app

import (
	"api-apotek/config"
	"api-apotek/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	os.Setenv("secret", "gin-apotek-bro0")
	db := config.DatabaseConfig.GetDatabase(config.DatabaseConfig{})
	config.MigrateDB(db)

	app := routes.Routes.GetRoutes(routes.Routes{})
	return app
}
