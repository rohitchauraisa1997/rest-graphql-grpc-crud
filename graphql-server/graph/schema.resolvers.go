package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/rohitchauraisa1997/graphql-server/graph/model"
	"github.com/rohitchauraisa1997/graphql-server/repository"
)

var videoRepo repository.VideoRepository = repository.NewVideoRepository()

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	video := &model.Video{
		ID:     strconv.Itoa(rand.Int()),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.User{ID: strconv.Itoa(rand.Int()), Name: input.Author},
	}
	r.videos = append(r.videos, video)

	videoRepo.Save(video)
	return video, nil
}

// DeleteVideo is the resolver for the deleteVideo field.
func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (*model.Video, error) {
	deletedVideo := videoRepo.DeleteVideo(id)
	return deletedVideo, nil
}

// UpdateVideo is the resolver for the updateVideo field.
func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input *model.NewVideo) (*model.Video, error) {
	updatedVideo := videoRepo.UpdateVideo(id, input)
	return updatedVideo, nil
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return videoRepo.FindAll(), nil
}

// GetVideo is the resolver for the getVideo field.
func (r *queryResolver) GetVideo(ctx context.Context, id string) (*model.Video, error) {
	video := videoRepo.FindVideo(id)
	return video, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
