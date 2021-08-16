package categories_controller

import (
	categories_query "gingonic/querys/categories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategoriesTrue(c *gin.Context) {		

	page := c.Query("page")

	data, err, pages := categories_query.ListCategoriesTrue(page)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot fetch the categories",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"totalPages":pages,
		"data":data,
	})
	return
}