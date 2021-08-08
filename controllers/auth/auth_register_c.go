package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Register",
	})
}
