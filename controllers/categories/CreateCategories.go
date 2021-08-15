package categories_controller

import (
	"fmt"
	"gingonic/core"
	categories_query "gingonic/querys/categories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateCategory(c *gin.Context){ 

	var body core.Category_Response

	if err := c.ShouldBindBodyWith(&body, binding.JSON);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"data invalid try it again",
		})
		return
	}

	//debug
	fmt.Print("[debug]")
	fmt.Println(body)
	
	ok, err := categories_query.CreateCategory(body)

	if !ok || err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"cannot create new category",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message":"created successfully",
	})
}