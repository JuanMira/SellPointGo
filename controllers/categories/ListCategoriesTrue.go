package categories_controller

import (
	categories_query "gingonic/querys/categories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategoriesTrue(c *gin.Context) {
	data, err := categories_query.ListCategoriesTrue()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot fetch the categories",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":data,
	})
	return
}