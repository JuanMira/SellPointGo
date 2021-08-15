package core

type Category struct {
	ID          uint64 `gorm:"primaryKey" json:"categoryId"`
	Name        string `gorm:"unique;not null" json:"categoryName"`
	Description string `json:"categoryDescription"`
	Status      bool   `gorm:"default:true"`
}

type Category_Response struct {
	ID          uint64 `json:"categoryId"`
	Name        string `json:"categoryName"`
	Description string `json:"categoryDescription"`
	Status      bool   `json:"categoryStatus,omitempty"`
}