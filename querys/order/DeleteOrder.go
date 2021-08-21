package order_query

import (
	"gingonic/core"
	"gingonic/services"
)

func DeleteQuery(id string) (bool, error) {

	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}
	data := core.Order{}
	db.Find(&data,id)
	data.Status = false
	err = db.Save(&data).Error

	if err != nil {
		return false, err
	}
	
	return true, nil
}