package routes

import (
	"go-absensi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Absensi API Running",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/hello", controllers.Hello)
		api.GET("/attendance", controllers.AttendanceIndex)
	}

}
