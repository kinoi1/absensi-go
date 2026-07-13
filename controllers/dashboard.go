package controllers

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Hello from Golang",
	})
}
