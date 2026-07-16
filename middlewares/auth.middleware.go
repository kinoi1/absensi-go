package middlewares

import (
	"go-absensi/utils" // Sesuaikan dengan path project Anda
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Header otorisasi diperlukan"})
			c.Abort() // Menghentikan request agar tidak lanjut ke handler/route
			return
		}

		// 2. Pastikan formatnya adalah "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format format Authorization harus 'Bearer <token>'"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. Validasi token JWT
		claims, err := utils.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau telah kedaluwarsa"})
			c.Abort()
			return
		}

		// 4. (Opsional pero Penting) Simpan data user dari token ke dalam Context Gin.
		// Ini berguna agar rute/handler di dalamnya tahu SIAPA yang sedang mengakses rute ini.
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		// 5. Token valid! Lanjutkan request ke handler tujuan
		c.Next()
	}
}
