package routes

import "gingonic/controllers/auth"

func authRoutes() {
	authRoute := router.Group("/api/auth")
	{
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/register", auth.Register_c)
	}
}