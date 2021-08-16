package categories_query

import (
	"gingonic/core"
	"gingonic/services"
	"math"
)

func ListCategories(page string) ([]core.Category,error, int64) {
	
	db,err := services.ConnectDB()	
	data := []core.Category{}
	var count int64
	if err != nil {
		return data, err, count
	}

	db.Find(&data)
	db.Model(&core.Category{}).Count(&count)
	
	totalPages := int(math.Ceil(float64(count)/10))

	return data, nil, int64(totalPages)
}