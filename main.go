package main

import (
    "go-absensi/config"
    "go-absensi/models"
    "go-absensi/routes"

    "github.com/gin-gonic/gin"
)

func main() {

    // Connect database
    config.ConnectDB()

    // Migration
    config.DB.AutoMigrate(
        &models.Users{},
        &models.Attendance{},
    )

    // Init Gin
    r := gin.Default()

    // Routes
    routes.SetupRoutes(r)

    // Default route
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Absensi API Running",
        })
    })

    // Run server
    r.Run(":8080")
}