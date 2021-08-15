package core

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID           uint64     `gorm:"primaryKey" json:"productId"`
	Name         string     `gorm:"not null" json:"productName"`
	Price        float64    `gorm:"not null" json:"productPrice"`
	Descriptions string     `json:"productDescription"`
	Image        string     `json:"productImage"`
	CategoryId   uint64     `json:"categoryId"`
	Category     []*Category `gorm:"many2many:product_categories;foreignKey:CategoryId"`
	Stock        int        `gorm:"not null;default:0" json:"productStock"`
}

//product response
type Product_R struct {
	Name         string  `json:"productName" form:"productName"`
	Price        float64 `json:"productPrice" form:"productPrice"`
	Descriptions string  `json:"productDescription" form:"productDescription"`
	Image        string  `json:"productImage" form:"productImage"`
	CategoryId   uint64  `json:"productCategoryId" form:"productCategoryId" `
	CreateDate   time.Time
	Stock        int `json:"productStock" form:"productStock"`
}

type Products_Response struct {
	ID           uint64  `json:"productId"`
	Name         string  `json:"productName"`
	Price        float64 `json:"productPrice"`
	Descriptions string  `json:"productDescription"`
	Image        string  `json:"productImage"`	
	CategoryName string  `json:"categoryName"`	
	Stock        int `json:"productStock"`
}
