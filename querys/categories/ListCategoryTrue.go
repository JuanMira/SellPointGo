package categories_query

import (
	"gingonic/core"
	"gingonic/services"
	"math"
	"strconv"
)

func ListCategoriesTrue(page string) ([]core.Category,error, int64){
	db, err := services.ConnectDB()

	var data []core.Category
	var count int64
	if err != nil{
		return data, err, count
	}	

	pageOffset,_ := strconv.Atoi(page)
	
	if pageOffset == 0{
		pageOffset = 1
	}
	
	offset := (pageOffset - 1) * 10	

	db.Where("status = ?", true).Offset(offset).Limit(10).Find(&data)
	db.Model(&core.Category{}).Where("status = ?",true).Count(&count)
	
	totalPages := int(math.Ceil(float64(count)/float64(10)))

	return data, nil, int64(totalPages)
}