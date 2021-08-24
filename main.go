package main

import (
	"fmt"
	"gingonic/migrations"
	"gingonic/routes"
	"gingonic/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	routes.Router = gin.Default()

	err := godotenv.Load()

	frontLink := os.Getenv("FRONT_LINK")

	routes.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"192.168.0.15",frontLink},
	}))

	

	gin.SetMode(gin.ReleaseMode)

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
