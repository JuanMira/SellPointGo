package detail_query

import (
	"gingonic/core"
	"gingonic/services"
)

func InsertDetail(data core.OrderDetail)(bool,error){
	db, err:= services.ConnectDB()

	if err != nil {
		return false, err		
	}

	result := db.Create(&data)

	if result.RowsAffected < 0 {
		return false, result.Error
	}	

	return true, nil
}