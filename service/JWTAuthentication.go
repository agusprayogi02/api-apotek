package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JWTAuth interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type AuthCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type JwtServices struct {
	SecretKey string
	Issuer    string
}

func JwtAuthService() JWTAuth {
	return &JwtServices{
		SecretKey: GetSecretKey(),
		Issuer:    "Agus Prayogi",
	}
}

func GetSecretKey() string {
	secret := os.Getenv("secret")
	if secret == "" {
		return "apotek"
	}
	return secret
}

func (service *JwtServices) GenerateToken(email string, isUser bool) string {
	claims := &AuthCustomClaims{email, isUser, jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(float64(time.Hour * 12)),
		Issuer:    service.Issuer,
		IssuedAt:  jwt.Now(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(service.SecretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(service.SecretKey), nil
	})
}
