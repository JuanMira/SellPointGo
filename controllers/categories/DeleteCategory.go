package categories_controller

import (
	categories_query "gingonic/querys/categories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	ok, err:= categories_query.DeleteCategory(id)

	if !ok || err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Cannot delete category",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":"Category deleted successfully",
	})
	return
}