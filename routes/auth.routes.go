package routes

import (
	"go-absensi/controllers"
	"go-absensi/middlewares" // Asumsi Anda punya middleware JWT di sini

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	// Rute Publik (Tidak butuh token untuk diakses)
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", controllers.LoginHandler)
		authGroup.POST("/refresh", controllers.RefreshHandler)
	}

	// Rute Terproteksi (Butuh Access Token yang valid)
	// Kita masukkan AuthMiddleware untuk menjaga rute di bawah ini
	protectedGroup := router.Group("/api/auth")
	protectedGroup.Use(middlewares.AuthMiddleware())
	{
		protectedGroup.POST("/logout", controllers.LogoutHandler)
	}
}
