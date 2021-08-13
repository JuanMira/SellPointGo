package core

type Country struct {
	ID          uint64 `gorm:"primaryKey" json:"countryId"`
	CountryName string `gorm:"not null" json:"countryName"`
}