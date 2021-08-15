package routes

import (
	product_controller "gingonic/controllers/Product"
	"gingonic/controllers/auth"
	"gingonic/middleware"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes() {

	//setup global middleware		

	authRoute := Router.Group("/api/auth")
	{
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/register", auth.Register_c)
	}

	//example
	storeRoute := Router.Group("/api/store")
	{
		storeRoute.GET("/", middleware.VerifyToken(),func(c *gin.Context) {
			c.JSON(200,gin.H{
				"debug":"test",
			})
		})
	}

	productRoute := Router.Group("/api/product")
	{
		productRoute.POST("/",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			product_controller.CreateProduct_c,
		)
		productRoute.GET("/all",product_controller.StoreList_c)		
		productRoute.GET("/")		
	}
}
