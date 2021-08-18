package product_query

import (
	"gingonic/core"
	"gingonic/services"
)

func UpdateProduct(id string, body core.Product_R) (bool, error) {

	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}

	var product core.Products

	db.First(&product,id)
		
	if body.Image == "" {
		product.Image = body.Image
	}

	product.Name = body.Name
	product.Descriptions = body.Descriptions
	product.CategoryId = body.CategoryId
	product.Price = body.Price
	product.Stock = body.Stock

	err = db.Save(&product).Error

	if err != nil {
		return false, err
	}

	return true, nil
}