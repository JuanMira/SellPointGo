package routes

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitializeRoutes() {

	//setup global middleware
	authRoutes()
	productRoutes()
	categoriesRoutes()
	orderRoutes()
}
