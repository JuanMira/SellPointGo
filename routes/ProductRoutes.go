package routes

import (
	product_controller "gingonic/controllers/Product"
	"gingonic/middleware"
)

func productRoutes() {
	productRoute := router.Group("/api/product")
	{
		productRoute.POST("/",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			product_controller.CreateProduct_c,
		)
		// all products true
		productRoute.GET("/all", product_controller.StoreList_c)
		// all products true/false
		productRoute.GET("/:category", product_controller.ListProductCategory)

		productRoute.DELETE("/:id",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			product_controller.DeleteProduct,
		)

		productRoute.PUT("/:id",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			product_controller.UpdateProduct,
		)
	}
}
