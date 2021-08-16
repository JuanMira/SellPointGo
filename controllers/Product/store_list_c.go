package product_controller

import (
	product_query "gingonic/querys/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreList_c(c *gin.Context) {

	page := c.Query("page")

	data, err, pages := product_query.ListProductPage(page)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"No data to fetch",
		})
		return
	}

	c.JSON(http.StatusAccepted,gin.H{
		"totalPages":pages,
		"data":data,
	})
	return
}