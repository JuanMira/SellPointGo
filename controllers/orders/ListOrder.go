package order_controller

import (
	order_query "gingonic/querys/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOrder(c *gin.Context){
	page := c.Query("page")
		

	data, totalPages,err := order_query.ListOrder(page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot fetch data try it again",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"totalPages":totalPages,
		"data":data,
	})
	return
}