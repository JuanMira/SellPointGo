package core

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID           uint64     `gorm:"primaryKey" json:"productId"`
	Name         string     `gorm:"not null" json:"productName"`
	Price        float64    `gorm:"not null" json:"productPrice"`
	Descriptions string     `json:"productDescription"`
	Image        string     `json:"productImage"`
	CategoryId  uint64 		`json:"categoryId"`
	Category   	[]Category 	`gorm:"many2many:product_categories;foreignKey:CategoryId"`
	Stock        int        `gorm:"not null;default:0" json:"productStock"`
}
