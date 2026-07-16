package models

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	RoleID    int       `json:"roleId"`
	RoleName  string    `json:"roleName"`
	NoHp      string    `json:"noHp"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
}
