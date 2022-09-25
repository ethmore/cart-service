package main

import (
	"cart-service/dotEnv"
	"cart-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{dotEnv.GoDotEnvVariable("BFF_URL")}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PublicRoutes(public)

	// private := router.Group("/")
	// routes.PrivateRoutes(private)

	if err := router.Run(":3007"); err != nil {
		panic(err)
	}
}
