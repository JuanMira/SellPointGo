package product_query

import (
	"gingonic/core"
	"gingonic/services"
	"math"
)

func ListProductPage(page string) ([]core.Products_Response,error, int64){

	db,err := services.ConnectDB()

	data := []core.Products_Response{}
	var count int64

	if err != nil {
		return data, err, count
	}					
		
	db.Model(&core.Products{}).Select(
		"products.id, products.name, products.price,products.descriptions,products.image,categories.name as category_name, products.stock",
	).Joins("JOIN categories ON categories.id = products.category_id").Where(
		"products.status = ?",
		true,
	).Scan(&data)

	db.Model(&core.Products{}).Count(&count)

	totalPages := int(math.Ceil(float64(count)/10))

	return data, nil, int64(totalPages)
}