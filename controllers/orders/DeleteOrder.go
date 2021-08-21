package order_controller

import (
	order_query "gingonic/querys/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOrder(c *gin.Context){
	id := c.Param("id")

	ok, err := order_query.DeleteQuery(id)

	if !ok || err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"that order isnt exist",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":"order deleted succesfully",
	})
	return
}