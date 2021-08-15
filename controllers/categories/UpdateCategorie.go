package categories_controller

import (
	"gingonic/core"
	categories_query "gingonic/querys/categories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func UpdateCategory(c *gin.Context){
	var bodyData core.Category_Response

	id := c.Param("id")

	if err := c.ShouldBindBodyWith(&bodyData, binding.JSON);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Data invalid please try it again",
		})
		return
	}

	categories_query.UpdateCategory(id,bodyData)

	c.JSON(http.StatusAccepted, gin.H{
		"message":"Category updated succesfully",
	})
	return
}