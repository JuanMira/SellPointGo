package core

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey"`
	Username  string    `gorm:"not null,uniqueIndex" json:"username"`
	Email     string   	`gorm:"not null,uniqueIndex" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Birthday  string 	`json:"birthday"`
	LastLogin time.Time	`json:"last_login"`
}


type UserBody_R struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Birthday string `json:"birthday"`
}
