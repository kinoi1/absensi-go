package controllers

import (
	"go-absensi/config"
	"go-absensi/models"
	"go-absensi/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler menangani endpoint /auth/login
func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[DEBUG LOG] Gagal bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format request tidak valid",
		})
		return
	}

	// Cari user berdasarkan email
	var user models.Users
	if err := config.DB.
		Where("email = ?", req.Username).
		First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}

	// Verifikasi password bcrypt
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}

	// Generate JWT
	accessToken, expiresIn, err := services.GenerateAccessToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat access token",
		})
		return
	}

	refreshToken, err := services.GenerateRefreshToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat refresh token",
		})
		return
	}

	response := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		Profile: models.UserProfile{
			ID:    string(rune(user.ID)),
			Name:  user.Name,
			Email: user.Email,
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

func LogoutHandler(c *gin.Context) {
	// 1. Ambil data user yang dimasukkan oleh middleware (jika Anda butuh mencatat log)
	username, exists := c.Get("username")
	if !exists {
		// Jika tidak ada di context, kita tetap lanjutkan proses hapus cookie
		username = "Unknown"
	}

	// 2. Hapus HttpOnly Cookie "refresh_token"
	// Cara menghapusnya adalah dengan mengeset MaxAge ke -1 dan value menjadi kosong ""
	c.SetCookie(
		"refresh_token", // Nama cookie yang sama saat login
		"",              // Kosongkan nilainya
		-1,              // Set -1 agar browser langsung menghapusnya seketika
		"/",             // Path harus sama dengan saat cookie dibuat
		"",              // Domain (kosongkan jika localhost, atau sesuaikan dengan domain Anda)
		true,            // Secure: true (wajib jika menggunakan HTTPS)
		true,            // HttpOnly: true (mencegah akses dari Javascript/XSS)
	)

	// 3. (Opsional) Logika tambahan ke Database / Redis
	// Jika Anda menyimpan whitelist/blacklist token di DB, hapus token milik user di sini:
	// db.DeleteRefreshTokenByUsername(username)

	// 4. Kembalikan respon sukses ke frontend
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Logout berhasil. Sesi Anda telah dihapus.",
		"user":    username,
	})
}
