package routes

import (
	order_controller "gingonic/controllers/orders"
	"gingonic/middleware"
)

func orderRoutes() {
	orderRoute := router.Group("/api/order")
	{
		orderRoute.POST(
			"/",
			middleware.VerifyToken(),
			order_controller.InsertOrder,
		)
	}
}