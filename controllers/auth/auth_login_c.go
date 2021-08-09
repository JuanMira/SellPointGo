package auth

import (
	auth_query "gingonic/querys/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var userBody *UserBody

	if err := c.ShouldBindBodyWith(&userBody, binding.JSON);err != nil {		
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"data invalid",
			"debug":userBody,
		})
		return		
	}	

	exist,err,user := auth_query.Login(userBody.Username,userBody.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})	
		return	
	}

	if exist == false{
		c.JSON(http.StatusNotFound,gin.H{
			"error":"User isn't exist",
		})
		return
		
	}

	c.JSON(http.StatusAccepted, gin.H{
		"messsage": "loged",
		"user":user,
	})
	return
}
