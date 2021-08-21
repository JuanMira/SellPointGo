package categories_query

import (
	"gingonic/core"
	"gingonic/services"
)

func UpdateCategory(idCategory string, bodyCategory core.Category_Response) (bool, error) {

	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}

	var category core.Category

	db.First(&category, idCategory)
	category.Name = bodyCategory.Name
	category.Description = bodyCategory.Description
	db.Save(&category)
	return true, nil
}