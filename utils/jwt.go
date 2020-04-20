package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("AllYourBase")

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	claim := Claims{
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: nowTime.Add(7 * time.Hour).Unix(),
			Issuer:    "gin-sam",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(jwtKey)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaim, e := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})

	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claims, nil
		}
	}
	return nil, e
}
