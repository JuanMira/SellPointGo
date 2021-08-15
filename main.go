package main

import (
	"fmt"
	"gingonic/migrations"
	"gingonic/routes"
	"gingonic/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	routes.Router = gin.Default()
	err := godotenv.Load()

	if err != nil {
		fmt.Println("file .env is missing")
		return
	} 

	sess := services.ConnectAWS()

	routes.Router.Use(func(c *gin.Context){
		c.Set("sess",sess)
		c.Next()
	})

	port := os.Getenv("PORT")

	migrations.Migrate()

	if port == "" {
		port = "5000"
	}
	//initialize routes
	routes.InitializeRoutes()
	routes.Router.Run(":" + port)
}
