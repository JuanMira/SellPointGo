package core

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              uint64    `gorm:"primaryKey" json:"orderId"`
	UserId          uint64    `gorm:"index" json:"userId,omitempty"`
	Ammount         int64     `json:"orderAmmount"`
	ShippingAddress string    `json:"orderShippingAddress"`
	OrderAddress    string    `json:"orderAddress"`
	OrderDate       time.Time `json:"orderDate,omitempty"`
	OrderStatus     bool      `gorm:"default:false" json:"orderStatus"`
	Status          bool      `gorm:"default:true" `	
}
