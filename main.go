package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"net/http"
)

func init() {
	viper.SetConfigFile("config/sonfig.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello Port :"+viper.GetString("port.router"))
	})
	e.Logger.Fatal(e.Start(":" + viper.GetString("port.router")))
}
