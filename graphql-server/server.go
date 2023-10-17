package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rohitchauraisa1997/graphql-server/http"
)

const defaultPort = "8082"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default()
	server.Use(cors.Default())
	// server.Use(middleware.BasicAuth())
	server.GET("/", http.PlayGroundHandler())
	server.POST("/query", http.GraphQLHandler())

	server.Run(":" + port)
}
