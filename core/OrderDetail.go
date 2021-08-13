package core

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	ID        uint64   `gorm:"primaryKey" json:"orderDetailId"`
	OrderId   uint64   `json:"orderId"`
	ProductId uint64   `json:"productId"`
	Price     float64  `json:"orderDetailPrice"`
	Quantity  int      `json:"orderDetailQuantity"`
	Order     Order    `gorm:"foreignKey:OrderId"`
	Product   Products `gorm:"foreignKey:ProductId"`
}
