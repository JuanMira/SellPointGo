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

	db.AutoMigrate(&core.User{})

}