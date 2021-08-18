package product_controller

import (
	"gingonic/core"
	product_query "gingonic/querys/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context){
	id := c.Param("id")

	var product core.Product_R

	if err := c.ShouldBind(&product);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"the data is invalid",
		})
		return
	}	

	ok, err := product_query.UpdateProduct(id, product)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot update the product pls try again",
		})
		return
	}

	c.JSON(http.StatusAccepted,gin.H{
		"message":"product update successfully",
	})
	return	
}