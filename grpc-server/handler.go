package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/rohitchauraisa1997/grpc-server/model"
	pb "github.com/rohitchauraisa1997/grpc-server/proto"
	"github.com/rohitchauraisa1997/grpc-server/repository"
)

var videoRepo repository.VideoRepository = repository.NewVideoRepository()

func (v *videoServer) GetVideos(context.Context, *pb.NoParam) (*pb.VideoList, error) {
	fmt.Println("GetVideos triggered")
	response := pb.VideoList{}
	videos := videoRepo.FindAll()
	fmt.Println(videos)
	var videoList []*pb.Video

	for _, video := range videos {
		responseVid := &pb.Video{}
		responseVid.Title = video.Title
		responseVid.Id = video.ID
		responseVid.Url = video.URL
		responseVid.Author = &pb.User{
			Id:   video.Author.ID,
			Name: video.Author.Name,
		}
		videoList = append(videoList, responseVid)
	}

	response.Videos = videoList

	return &response, nil
}

func (v *videoServer) GetVideo(ctx context.Context, id *pb.Id) (*pb.Video, error) {
	fmt.Println("GetVideo triggered for id = ", id.Id)
	video := videoRepo.FindVideo(id.Id)
	response := pb.Video{
		Id:    video.ID,
		Title: video.Title,
		Url:   video.URL,
	}
	response.Author = &pb.User{
		Id:   video.Author.ID,
		Name: video.Author.Name,
	}

	return &response, nil
}

func (v *videoServer) AddVideo(ctx context.Context, inputVideo *pb.NewVideo) (*pb.Video, error) {
	fmt.Println("Add triggered")
	videoToAdd := model.Video{
		ID:    strconv.Itoa(rand.Int()),
		Title: inputVideo.Title,
		URL:   inputVideo.Url,
	}
	videoToAdd.Author = &model.User{
		ID:   strconv.Itoa(rand.Int()),
		Name: inputVideo.Author,
	}

	videoRepo.Save(&videoToAdd)

	response := pb.Video{
		Id:    videoToAdd.ID,
		Title: videoToAdd.Title,
		Url:   videoToAdd.URL,
	}
	response.Author = &pb.User{
		Id:   videoToAdd.Author.ID,
		Name: videoToAdd.Author.Name,
	}

	return &response, nil
}

func (v *videoServer) DeleteVideo(ctx context.Context, id *pb.Id) (*pb.Video, error) {
	fmt.Println("DeleteVideo triggered for id = ", id.Id)
	video := videoRepo.DeleteVideo(id.Id)
	response := pb.Video{
		Id:    video.ID,
		Title: video.Title,
		Url:   video.URL,
	}
	response.Author = &pb.User{
		Id:   video.Author.ID,
		Name: video.Author.Name,
	}

	return &response, nil
}
