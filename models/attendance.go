package models

import "time"

type Attendance struct {
    ID        uint `gorm:"primaryKey"`
    UserID    uint
    CheckIn   time.Time
    CheckOut  time.Time
}