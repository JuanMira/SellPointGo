package order_controller

import (
	"gingonic/core"
	order_query "gingonic/querys/order"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func UpdateController(c *gin.Context){
	var orderResponse core.Order

	id := c.Param("id")

	if err:= c.ShouldBindBodyWith(&orderResponse, binding.JSON);err != nil{
		

		c.JSON(http.StatusBadGateway, gin.H{
			"error":"data invalid pls try again",
		})
		return
	}

	ok, err := order_query.UpdateOrder(id, orderResponse)

	if !ok || err != nil {		
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot update data pls ty again",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":"Order update succesfully",
	})
	return
}