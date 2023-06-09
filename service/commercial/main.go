package main

import (
	"commercial-propfloor/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.PrivateRoutes(router)

	log.Fatal(router.Run(":" + port))

}
