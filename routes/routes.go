package routes

import (
	"gingonic/controllers/auth"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes() {

	authRoute := Router.Group("/api/auth")
	{
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/register", auth.Register)
	}
}
