package seeders

import (
	"go-absensi/config"
	"go-absensi/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	var count int64
	config.DB.Model(&models.Users{}).Count(&count)

	if count > 0 {
		return
	}
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte("123456"),
		bcrypt.DefaultCost,
	)

	user := models.Users{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: string(hash),
	}

	config.DB.FirstOrCreate(
		&user,
		models.Users{Email: user.Email},
	)
}
