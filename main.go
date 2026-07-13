package main

import (
	"go-absensi/config"
	"go-absensi/routes"
	"go-absensi/seeders"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect database
	config.ConnectDB()
	config.MigrateDB()
	seeders.SeedAttendance()
	seeders.SeedUsers()
	// Init Gin
	r := gin.Default()

	r.Use(cors.New(config.CORSConfig()))
	// Routes
	routes.SetupRoutes(r)

	// Run server
	r.Run(":8080")
}
