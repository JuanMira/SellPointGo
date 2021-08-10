package routes

import (
	"gingonic/controllers/auth"
	"gingonic/middleware"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes() {

	authRoute := Router.Group("/api/auth")
	{
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/register", auth.Register_c)
	}

	storeRoute := Router.Group("/api/store")
	{
		storeRoute.GET("/", middleware.VerifyToken(),func(c *gin.Context) {
			c.JSON(200,gin.H{
				"debug":"test",
			})
		})
	}
}
