package routes

import (
	product_controller "gingonic/controllers/Product"
	"gingonic/middleware"
)

func productRoutes() {
	productRoute := Router.Group("/api/product")
	{
		productRoute.POST("/",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			product_controller.CreateProduct_c,
		)
		// all products true
		productRoute.GET("/all", product_controller.StoreList_c)
		// all products true/false
		productRoute.GET("/")
	}
}