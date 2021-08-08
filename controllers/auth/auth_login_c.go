package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"messsage": "loged",
	})
}
