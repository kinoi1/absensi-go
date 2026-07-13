package seeders

import (
	"go-absensi/config"
	"go-absensi/models"
	"math/rand"
	"time"
)

func SeedAttendance() {
	var count int64
	config.DB.Model(&models.Attendance{}).Count(&count)

	if count > 0 {
		return
	}

	var attendances []models.Attendance

	for i := 1; i <= 20; i++ {
		// tanggal 20 hari terakhir
		date := time.Now().AddDate(0, 0, -i)

		// jam masuk antara 07:30 - 08:30
		checkIn := time.Date(
			date.Year(),
			date.Month(),
			date.Day(),
			7+rand.Intn(2),
			rand.Intn(60),
			0,
			0,
			time.Local,
		)

		// jam pulang antara 16:00 - 17:30
		checkOut := time.Date(
			date.Year(),
			date.Month(),
			date.Day(),
			16+rand.Intn(2),
			rand.Intn(60),
			0,
			0,
			time.Local,
		)

		statuses := []string{
			"present",
			"late",
			"leave",
			"absent",
		}

		attendance := models.Attendance{
			UserID:   uint(rand.Intn(5) + 1), // User ID 1-5
			Date:     date,
			CheckIn:  checkIn,
			CheckOut: checkOut,
			Status:   statuses[rand.Intn(len(statuses))],
		}

		attendances = append(attendances, attendance)
	}

	if err := config.DB.Create(&attendances).Error; err != nil {
		panic(err)
	}
}
