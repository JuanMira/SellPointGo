package categories_query

import (
	"gingonic/core"
	"gingonic/services"
)

func ListCategoriesTrue() ([]core.Category,error){
	db, err := services.ConnectDB()

	var data []core.Category

	if err != nil{
		return data, err
	}	

	db.Debug().Where("status = ?", true).Find(&data)

	return data, nil
}