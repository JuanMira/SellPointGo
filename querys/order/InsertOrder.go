package order_query

import (
	"gingonic/core"
	"gingonic/services"
	"time"
)

func InsertOrder(data core.Order) ( bool,error){ 

	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}
	
	data.OrderDate = time.Now()
	data.OrderStatus = false

	result := db.Create(&data)

	if result.RowsAffected < 0 {
		return false, result.Error
	}

	return true, nil
}