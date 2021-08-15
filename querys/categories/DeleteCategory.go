package categories_query

import (
	"gingonic/core"
	"gingonic/services"
)

func DeleteCategory(id string) (bool, error) {
	db, err := services.ConnectDB()

	if err != nil {
		return false, err
	}

	var category core.Category

	db.First(&category, id)
	category.Status = false
	db.Save(&category)
	return true, nil
}