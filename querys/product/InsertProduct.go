package product_query

import (
	"fmt"
	"gingonic/core"
	"gingonic/services"
)

func CreateProduct(product core.Product_R) (bool,error) {

	db, err := services.ConnectDB()

	if err != nil {
		return false,err
	}

	var product_db core.Products = core.Products{
		Name: product.Name,
		Descriptions: product.Descriptions,
		Image: product.Image,
		CategoryId: product.CategoryId,
		Price: product.Price,
		Stock: product.Stock,		
	}

	result := db.Create(&product_db)

	fmt.Print("Rows affected = ")
	fmt.Println(result.RowsAffected)
	return true, nil
}