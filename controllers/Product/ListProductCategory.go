package product_controller

import (
	product_query "gingonic/querys/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductCategory(c *gin.Context) {
	categoryId := c.Param("category")

	if categoryId == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Param category is required",
		})
		return
	}

	data, err := product_query.ListProductCategory(categoryId)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"no data to fetch",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":data,
	})
	return
}