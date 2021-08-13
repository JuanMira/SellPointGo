package core

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint64    `gorm:"primaryKey"`
	Username       string    `gorm:"not null,uniqueIndex" json:"username"`
	Email          string    `gorm:"not null,uniqueIndex" json:"email"`
	Password       string    `gorm:"not null" json:"password"`
	FirstName      string    `gorm:"not null" json:"firstName"`
	LastName       string    `gorm:"not null" json:"lastName"`
	Phone          string    `gorm:"not null" json:"phone"`
	BillingAddress string    `gorm:"not null" json:"billingAddress"`
	CountryId      uint64    `json:"countryId"`
	Birthday       string    `json:"birthday"`
	LastLogin      time.Time `json:"last_login"`
	Country        Country   `gorm:"foreignKey:CountryId"`
	Role           bool      `gorm:"default:0"`
}

type UserBody_R struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	CountryId      uint64 `json:"countryId"`
	Phone          string `json:"phone"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	BillingAddress string `json:"billingAddress"`
}
