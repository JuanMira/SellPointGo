package detail_query

import (
	"gingonic/core"
	"gingonic/services"
)

func DeleteDetail(id string) (bool, error) {

	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}

	var data core.OrderDetail

	db.Find(&data, id)
	data.Status = false
	err = db.Save(&data).Error

	if err != nil {
		return false, err
	}

	return true, nil
}