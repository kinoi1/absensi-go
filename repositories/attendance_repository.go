package repositories

import (
	"go-absensi/config"
	"go-absensi/models"
)

func GetAttendances(page, limit int) ([]models.Attendance, int64, error) {
	var attendances []models.Attendance
	var total int64

	// Hitung total data
	if err := config.DB.
		Model(&models.Attendance{}).
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	// Ambil data sesuai pagination
	err := config.DB.
		Order("date DESC").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&attendances).
		Error

	if err != nil {
		return nil, 0, err
	}

	return attendances, total, nil
}
