package handlers

import (
    "time"

    "go-absensi/config"
    "go-absensi/models"

    "github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
    attendance := models.Attendance{
        UserID:  1,
        CheckIn: time.Now(),
    }

    config.DB.Create(&attendance)

    c.JSON(200, gin.H{
        "message": "Check in berhasil",
    })
}