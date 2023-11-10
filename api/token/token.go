package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("37248dc2d0572d57d8230e8e495fa5b179d819167767f4678ea78bef065b5c75")

type Claims struct {
	jwt.StandardClaims
}

func GenerateRefreshToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString(JwtKey)
}

func ValidateToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	return token, claims, err
}
