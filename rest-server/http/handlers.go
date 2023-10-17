package http

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohitchauraisa1997/rest-server/model"
	"github.com/rohitchauraisa1997/rest-server/repository"
)

var videoRepo repository.VideoRepository = repository.NewVideoRepository()

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "ping test for rest server!!")
}

func AddVideo(c *gin.Context) {
	fmt.Println("Add Video triggered")
	var video model.Video

	if err := c.BindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video.ID = strconv.Itoa(rand.Int())
	video.Author.ID = strconv.Itoa(rand.Int())
	videoRepo.Save(&video)

	c.JSON(http.StatusOK, video)
}

func GetAllVideos(c *gin.Context) {
	fmt.Println("GetAllVideos triggered")
	videos := videoRepo.FindAll()
	c.JSON(http.StatusOK, videos)
}

func GetVideoById(c *gin.Context) {
	fmt.Println("GetVideoById triggered")
	entryId := c.Query("id")
	video := videoRepo.FindVideo(entryId)
	c.JSON(http.StatusOK, video)
}

func DeleteVideo(c *gin.Context) {
	fmt.Println("DeleteVideo triggered")
	entryId := c.Query("id")
	video := videoRepo.DeleteVideo(entryId)
	c.JSON(http.StatusOK, video)
}

func UpdateVideo(c *gin.Context) {
	fmt.Println("UpdateVideo triggered")
	entryId := c.Query("id")

	var newVideo model.Video

	if err := c.BindJSON(&newVideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video := videoRepo.UpdateVideo(entryId, &newVideo)

	c.JSON(http.StatusOK, video)

}
