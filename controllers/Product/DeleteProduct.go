package product_controller

import (
	product_query "gingonic/querys/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context){
	id := c.Param("id")

	ok, err := product_query.DeleteProduct(id)


	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot delete thata product pls try again",
		})
		return
	}
	
	c.JSON(http.StatusAccepted, gin.H{
		"message":"Product deleted succesfully",
	})
	return
}