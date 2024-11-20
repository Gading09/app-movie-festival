package model

import "time"

type User struct {
	Id        string    `gorm:"primaryKey"`
	Username  string    `gorm:"size:100;not null;unique"`
	Password  string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:100;unique"`
	IsAdmin   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}
