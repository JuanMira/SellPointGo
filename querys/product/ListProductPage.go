package product_query

import (
	"gingonic/core"
	"gingonic/services"
)

func ListProductPage() ([]core.Products_Response,error){

	db,err := services.ConnectDB()

	data := []core.Products_Response{}

	if err != nil {
		return data, err
	}					
		
	db.Debug().Model(&core.Products{}).Select(
		"products.id, products.name, products.price,products.descriptions,products.image,categories.name as category_name, products.stock",
	).Joins("JOIN categories ON categories.id = products.category_id").Scan(&data)


	return data, nil
}