package order_query

import (
	"gingonic/core"
	"gingonic/services"
	"math"
	"strconv"
)

func ListOrder(page string) ([]core.Order,int64,error){
	var data []core.Order
	var count int64
	db, err := services.ConnectDB()

	if err != nil{
		return data, 0,err
	}

	pageOffset, _ := strconv.Atoi(page)

	if pageOffset == 0{ 
		pageOffset = 1
	}

	offset := (pageOffset - 1 ) * 10

	db.Where("status = ?",true).Offset(offset).Limit(10).Find(&data)
	db.Model(&core.Order{}).Count(&count)

	totalPages := int(math.Ceil(float64(count)/10))

	return data, int64(totalPages) ,nil
}