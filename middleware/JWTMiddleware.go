package middleware

import (
	"api-apotek/service"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := service.JwtAuthService().ValidateToken(tokenString)
		if !token.Valid {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		claims := token.Claims.(*jwt.MapClaims)
		fmt.Println(claims)
	}
}
