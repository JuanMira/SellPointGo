package product_query

import (
	"gingonic/core"
	"gingonic/services"
	"math"
	"strconv"
)

func ListProductCategory(categoryId string, page string) ([]core.Products_Response,error, int64){
	db,err := services.ConnectDB()
	
	data := []core.Products_Response{}
	var count int64 

	if err != nil {
		return data, err, count
	}

	pageOffset,_ := strconv.Atoi(page)

	if pageOffset == 0{
		pageOffset = 1
	}

	offset := (pageOffset - 1 ) * 10

	db.Model(&core.Products{}).Select(
		"products.id, products.name, products.price,products.descriptions,products.image,categories.name as category_name, products.stock",
	).Joins("JOIN categories ON categories.id = products.category_id").Where(
		"products.category_id = ? AND products.status = ?",
		categoryId,
		true,
	).Offset(offset).Limit(10).Scan(&data)

	db.Model(&core.Products{}).Joins(
		"JOIN categories ON categories.id = products.category_id",
	).Count(&count)



	totalPages := int(math.Ceil(float64(count)/float64(10)))

	return data, nil, int64(totalPages)
}