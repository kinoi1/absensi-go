package routes

import (
    "go-absensi/handlers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/checkin", handlers.CheckIn)
}