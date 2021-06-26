package middleware

import (
	"api-apotek/service"
	"api-apotek/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		str := utils.ExtactToken(authHeader)
		// fmt.Println(str)
		if str == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, err := service.JwtAuthService().ValidateToken(str)
		if !token.Valid {
			fmt.Println("Error: ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// claims := token.Claims.(jwt.MapClaims)
		// fmt.Println(claims)
		c.Next()
	}
}
