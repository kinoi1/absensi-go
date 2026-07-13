package models

import (
	"time"
)

type Attendance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userID" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"type:date;null"`
	CheckIn   time.Time `json:"checkOut"`
	CheckOut  time.Time `json:"checkIn"`
	Status    string    `json:"status" gorm:"type:varchar(20);default:'present'"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
