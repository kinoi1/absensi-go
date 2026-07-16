package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// Ganti dengan secret key Anda yang aman (sebaiknya dari env variable)
var JWTSecret = []byte("super_secret_key_anda_123")

type JWTClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// ValidateAccessToken memvalidasi token dan mengembalikan data claims di dalamnya
func ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing-nya HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metode signing tidak terduga")
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}
