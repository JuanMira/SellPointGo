package main

import (
	"gingonic/migrations"
	"gingonic/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	routes.Router = gin.Default()
	port := os.Getenv("PORT")

	migrations.Migrate()

	if port == "" {
		port = "5000"
	}
	//initialize routes
	routes.InitializeRoutes()
	routes.Router.Run(":" + port)
}
