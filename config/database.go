package config

import (
	"go-absensi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123456 dbname=absensi port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.Attendance{},
		&models.Users{},
	)
	if err != nil {
		panic(err)
	}
}
