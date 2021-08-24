package routes

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes() {	

	//setup global middleware
	authRoutes()
	productRoutes()
	categoriesRoutes()
	orderRoutes()
	orderDetailRoute()
}
