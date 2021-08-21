package routes

import (
	categories_controller "gingonic/controllers/categories"
	"gingonic/middleware"
)

func categoriesRoutes() {
	categoriesRoute := router.Group("/api/category")
	{
		categoriesRoute.GET("/",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			categories_controller.ListCategories,
		)

		categoriesRoute.POST("/",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			categories_controller.CreateCategory,
		)

		categoriesRoute.PUT("/:id",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			categories_controller.UpdateCategory,
		)

		categoriesRoute.DELETE("/:id",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			categories_controller.DeleteCategory,
		)

		categoriesRoute.GET("/all",
			middleware.VerifyToken(),
			middleware.VerifyAdmin(),
			categories_controller.ListCategoriesTrue,
		)
	}

}