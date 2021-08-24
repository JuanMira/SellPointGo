package detail_query

import (
	"gingonic/core"
	"gingonic/services"
)

type data struct{
	ID string
	Name string
	Price float64
	Quantity int
	Ammount int64
	OrderStatus bool	
}

func ListDetail() ([]data,error){
	db, err := services.ConnectDB()

	var data []data

	if err != nil {
		return data, nil
	}

	selects := "order_details.id,order_details.price, order_details.quantity,products.name, orders.ammount, orders.status"

	err = db.Model(&core.OrderDetail{}).Select(
		selects,	
	).Joins(
		"JOIN products ON products.id = order_details.product_id",
	).Joins(
		"JOIN orders ON orders.id = order_details.order_id",
	).Where(
		"order_details.status = ?",
		true,
	).Scan(&data).Error	

	if err != nil{
		return data, err
	}		

	return data, nil
}	