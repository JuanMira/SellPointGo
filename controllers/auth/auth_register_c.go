package auth

import (
	"fmt"
	"gingonic/core"
	auth_query "gingonic/querys/auth"
	"gingonic/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Register_c(c *gin.Context) {

	var userBody core.UserBody_R

	if err := c.ShouldBindBodyWith(&userBody, binding.JSON);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"data invalid",
			"message-debug":"data json not parsed",
		})	
		fmt.Println(userBody)
		return 	
	}

	cryptPass, err := services.HashPassword(userBody.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return
	}

	userBody.Password = cryptPass

	work,err := auth_query.Register(userBody)

	if work == false || err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Something was wrong try again",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "user register successfully",
		"status":"created",
	})
}