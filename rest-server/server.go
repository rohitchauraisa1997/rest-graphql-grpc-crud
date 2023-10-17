package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rohitchauraisa1997/rest-server/http"
)

func main() {
	fmt.Println("Running rest serverr")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	router.GET("/ping", http.Ping)
	router.POST("/video/add", http.AddVideo)
	router.GET("/videos/all", http.GetAllVideos)
	router.GET("/video", http.GetVideoById)
	router.PUT("/video/update", http.UpdateVideo)
	router.DELETE("/video/delete", http.DeleteVideo)

	router.Run(":" + port)
}
