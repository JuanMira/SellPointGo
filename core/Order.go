package core

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              uint64    `gorm:"primaryKey" json:"orderId"`
	UserId          uint64    `json:"userId"`
	Ammount         int64     `json:"orderAmmount"`
	ShippingAddress string    `json:"orderShippingAddress"`
	OrderAddress    string    `json:"orderAddress"`
	OrderDate       time.Time `json:"orderDate"`
	OrderStatus     bool      `gorm:"default:0"`
	User            User      `gorm:"foreignKey:UserId"`
}
