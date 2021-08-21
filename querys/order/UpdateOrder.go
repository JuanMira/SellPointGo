package order_query

import (
	"gingonic/core"
	"gingonic/services"
)

func UpdateOrder(id string, data core.Order) (bool, error) {	

	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}

	var order core.Order

	db.First(&order, id)
	order.ShippingAddress = data.ShippingAddress
	order.OrderStatus = data.OrderStatus
	err = db.Save(&order).Error

	if err != nil {
		return false, err
	}

	return true, nil
}