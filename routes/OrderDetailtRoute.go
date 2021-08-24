package routes

import (
	order_controller "gingonic/controllers/orders"
	"gingonic/middleware"
)

func orderDetailRoute() {
	orderDetailRoute := Router.Group("/api/order/detail")
	{
		orderDetailRoute.POST(
			"/",
			middleware.VerifyToken(),
			order_controller.InsertDetail,
		)

		orderDetailRoute.GET(
			"/",
			middleware.VerifyToken(),
			order_controller.ListDetail,
		)

		orderDetailRoute.DELETE(
			"/:id",
			middleware.VerifyToken(),
			order_controller.DeleteDetail,			
		)
	}
}