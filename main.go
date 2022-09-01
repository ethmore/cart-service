package main

import (
	"cart/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PublicRoutes(public)

	// private := router.Group("/")
	// routes.PrivateRoutes(private)

	router.Run(":3007")
}
