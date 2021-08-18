package product_query

import (
	"gingonic/core"
	"gingonic/services"
)

func DeleteProduct(id string) (bool, error) {
	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}

	var product *core.Products

	db.First(&product,id)
	product.Status = false
	db.Save(&product)	
	return true, nil
}