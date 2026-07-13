package controllers

import (
	"go-absensi/repositories"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AttendanceIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))

	attendances, total, err := repositories.GetAttendances(page, limit)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed to get attendance data",
			"error":   err.Error(),
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(200, gin.H{
		"success": true,
		"message": "Attendance List",
		"data":    attendances,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total_data":   total,
			"total_pages":  totalPages,
			"has_next":     page < totalPages,
			"has_prev":     page > 1,
		},
	})
}
