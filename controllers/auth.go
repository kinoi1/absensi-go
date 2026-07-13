package controllers

import (
	"go-absensi/models"
	"go-absensi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler menangani endpoint /auth/login
func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Format request tidak valid"})
		return
	}

	// MOCK VALIDASI: Silakan ganti bagian ini dengan query database asli Anda
	if req.Username != "admin" || req.Password != "password123" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	// Jika kredensial cocok, buatkan JWT pasangannya
	accessToken, expiresIn, err := services.GenerateAccessToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat access token"})
		return
	}

	refreshToken, err := services.GenerateRefreshToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat refresh token"})
		return
	}

	// Susun response sesuai kebutuhan NextAuth di frontend
	response := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		Profile: models.UserProfile{
			ID:    "USR-001",
			Name:  "Temi The Creator",
			Email: "temi@example.com",
		},
	}

	c.JSON(http.StatusOK, response)
}

// RefreshHandler menangani endpoint /auth/refresh
func RefreshHandler(c *gin.Context) {
	var req models.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Refresh token wajib diisi"})
		return
	}

	// Validasi token refresh yang dikirim dari Next.js
	username, err := services.ValidateToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token tidak valid atau kedaluwarsa"})
		return
	}

	// Jika valid, terbitkan Access Token baru yang segar
	newAccessToken, expiresIn, err := services.GenerateAccessToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui access token"})
		return
	}

	// Kirimkan balik ke Next.js
	response := models.RefreshResponse{
		AccessToken: newAccessToken,
		ExpiresIn:   expiresIn,
	}

	c.JSON(http.StatusOK, response)
}
