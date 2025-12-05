package utils

import (

	"study/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


var jwtSecret = []byte("123456789")

func GenerateToken(user *models.User) (string, error){
	claims := &models.Claims{
		UserID: user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}


func ParseToken(tokenString string) (*models.Claims, error){
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid{
		return nil, err
	}
	return claims, nil
}