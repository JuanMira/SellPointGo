package core

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    uint      `gorm:"primaryKey"`
	Username  string    `gorm:"not null,uniqueIndex"`
	Email     *string   `gorm:"not null,uniqueIndex"`
	Password  string    `gorm:"not null"`
	Birthday  time.Time 
	LastLogin time.Time	
}