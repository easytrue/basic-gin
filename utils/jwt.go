package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var secretKey = []byte(viper.GetString("jwt.secretKey"))

type jwtAuthClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id uint, name string) (string, error) {
	authClaims := jwtAuthClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.expiresAt") * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*jwtAuthClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwtAuthClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		err = errors.New("无权访问")
	}

	if !token.Valid {
		err = errors.New("Invalid Token")
	}

	return token.Claims.(*jwtAuthClaims), err
}
