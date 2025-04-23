package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var app_key = []byte(os.Getenv("APP_KEY"))

type JWTClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	return token.SignedString(app_key)
}

func VerifyToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return app_key, nil
	})

	if err != nil || !token.Valid {
		log.Println(err)
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		log.Println("Invalid token claims")
		return nil, err
	}

	return claims, nil
}
