package order_controller

import (
	"fmt"
	"gingonic/core"
	"net/http"

	detail_query "gingonic/querys/orderDetail"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InsertDetail(c *gin.Context) {
	var orderResponse core.OrderDetail

	if err := c.ShouldBindBodyWith(&orderResponse, binding.JSON);err != nil{
		//debug
		fmt.Println(orderResponse)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid data pls try again",
		})
		return
	}

	//debug
	fmt.Println(orderResponse)

	ok,err := detail_query.InsertDetail(orderResponse)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"cannot insert data try again",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":"orderDetail inserted succesfully",
	})
	return
}