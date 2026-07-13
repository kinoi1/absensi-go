package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Gunakan secret key yang aman (sebaiknya simpan di .env)
var jwtSecret = []byte("super-secret-key-anda-12345")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateAccessToken membuat token berdurasi pendek (15 menit)
func GenerateAccessToken(username string) (string, int64, error) {
	duration := 15 * time.Minute
	expiresAt := time.Now().Add(duration)

	claims := &JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, int64(duration.Seconds()), err
}

// GenerateRefreshToken membuat token berdurasi panjang (7 hari)
func GenerateRefreshToken(username string) (string, error) {
	claims := &JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken memvalidasi string JWT dan mengembalikan username di dalamnya
func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", errors.New("invalid token claims")
}
