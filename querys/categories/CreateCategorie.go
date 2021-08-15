package categories_query

import (
	"fmt"
	"gingonic/core"
	"gingonic/services"
)

func CreateCategory(data core.Category_Response) (bool, error) {
	db, err := services.ConnectDB()

	cData := core.Category{
		Name: data.Name,
		Description: data.Description,	
	}

	if err != nil{
		return false, err
	}

	rows := db.Create(&cData)

	fmt.Print("Rows affected = ")
	fmt.Println(rows.RowsAffected)

	return true, nil
}