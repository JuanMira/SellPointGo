package order_controller

import (
	detail_query "gingonic/querys/orderDetail"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListDetail(c *gin.Context){

	data, err := detail_query.ListDetail()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot fetch data ",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":"data fetched successfully",
		"data":data,
	})
	return
}