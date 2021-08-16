package product_controller

import (
	product_query "gingonic/querys/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductCategory(c *gin.Context) {
	categoryId := c.Param("category")
	page := c.Query("page")
	if categoryId == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Param category is required",
		})
		return
	}

	// message 
	data, err, pages := product_query.ListProductCategory(categoryId, page)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"no data to fetch",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"totalPages":pages,
		"data":data,
	})
	return
}