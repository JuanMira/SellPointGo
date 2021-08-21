package order_controller

import (
	"gingonic/core"
	order_query "gingonic/querys/order"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InsertOrder(c *gin.Context) {
	var orderResponse core.Order

	if err := c.ShouldBindBodyWith(&orderResponse, binding.JSON);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"data invalid try it again",
		})
		
		return
	}	

	ok, err := order_query.InsertOrder(orderResponse)

	if !ok || err != nil {		
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"the data cannot be saved try again",
		})
		return
	}

	c.JSON(http.StatusAccepted,gin.H{
		"message":"order inserted succesfully",
	})
	return
}