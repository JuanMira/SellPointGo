package order_controller

import (
	detail_query "gingonic/querys/orderDetail"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDetail(c *gin.Context){
	id := c.Param("id")

	ok, err :=   detail_query.DeleteDetail(id)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot delete the detail",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":"detail deleted succesfully",
	})
	return
}