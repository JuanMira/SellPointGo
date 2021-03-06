package migrations

import (
	"gingonic/core"
	"gingonic/services"
)

func Migrate() {
	db,err := services.ConnectDB()

	if err != nil {
		return
	}	

	db.AutoMigrate(
		&core.User{},
		&core.Country{},
		&core.Products{},
		&core.Category{},
		&core.Order{},
		&core.OrderDetail{},
	)

}