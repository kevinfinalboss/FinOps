package middlewares

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("37248dc2d0572d57d8230e8e495fa5b179d819167767f4678ea78bef065b5c75")

type Claims struct {
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, err
}
