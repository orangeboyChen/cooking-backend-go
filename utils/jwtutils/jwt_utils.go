package jwtutils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

const key = "Ho0k6GW7aFUcntjs"

func CreateJwtToken(userId string) (string, error) {
	raw := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"timestamp": time.Now().UnixMilli(),
	})

	var token string
	token, err := raw.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeJwtToken(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	return claim.Claims.(jwt.MapClaims), nil
}
