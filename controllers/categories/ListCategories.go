package categories_controller

import (
	categories_query "gingonic/querys/categories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context){
	data,err := categories_query.ListCategories()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot fetch all data",			
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":data,
	})
}