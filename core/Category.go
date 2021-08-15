package core

type Category struct {
	ID          uint64    `gorm:"primaryKey" json:"categoryId"`
	Name        string    `gorm:"unique;not null" json:"categoryName"`
	Description string    `json:"categoryDescription"`
	Products    *Products `gorm:"many2many:product_categories"`
}