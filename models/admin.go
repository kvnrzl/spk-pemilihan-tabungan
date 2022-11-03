package model

import "time"

type Admin struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"username"`
	Password  string `json:"password"`
}
