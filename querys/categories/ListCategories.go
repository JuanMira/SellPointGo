package categories_query

import (
	"gingonic/core"
	"gingonic/services"
)

func ListCategories() ([]core.Category,error) {
	data := []core.Category{}
	db,err := services.ConnectDB()

	if err != nil {
		return data, err
	}

	db.Find(&data)
	
	return data, nil
}