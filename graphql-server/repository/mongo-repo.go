package repository

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/rohitchauraisa1997/graphql-server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE = "videodb"
	COLL     = "video-coll"
)

type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []*model.Video
	FindVideo(id string) *model.Video
	DeleteVideo(id string) *model.Video
	UpdateVideo(id string, updatedVideo *model.NewVideo) *model.Video
}

type database struct {
	Client *mongo.Client
}

func NewVideoRepository() VideoRepository {
	// The NewVideoRepository() function creates an instance of the database struct,
	//  which implements the VideoRepository interface.
	// Since Go allows you to return a pointer to a struct that implements an interface,
	//  you can return &database{} as a VideoRepository, making it possible to use this function
	// to obtain an object that conforms to the VideoRepository interface.

	// when db running in docker compose but server in local.
	// MONGODB := "mongodb://video-user:video-pwd@localhost:27023/?authSource=videodb"

	// when db and server running both in docker compose.
	MONGODB := "mongodb://video-user:video-pwd@mongodb:27017/?authSource=videodb"
	clientOptions := options.Client().ApplyURI(MONGODB)
	clientOptions = clientOptions.SetMaxPoolSize(500)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &database{
		Client: dbClient,
	}
}

func (db *database) Save(video *model.Video) {
	collection := db.Client.Database(DATABASE).Collection(COLL)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err)
	}

}

func (db *database) FindAll() []*model.Video {
	collection := db.Client.Database(DATABASE).Collection(COLL)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var result []*model.Video
	for cursor.Next(context.TODO()) {
		var v *model.Video
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}

	return result
}

func (db *database) FindVideo(objectID string) *model.Video {
	collection := db.Client.Database(DATABASE).Collection(COLL)
	var result model.Video
	err := collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal("Error finding document:", err)
	}
	return &result
}

func (db *database) DeleteVideo(objectID string) *model.Video {
	collection := db.Client.Database(DATABASE).Collection(COLL)
	var result model.Video
	err := collection.FindOneAndDelete(context.TODO(), bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal("Error finding document:", err)
	}
	return &result
}

func (db *database) UpdateVideo(id string, newInputVideo *model.NewVideo) *model.Video {
	collection := db.Client.Database(DATABASE).Collection(COLL)
	filter := bson.M{"_id": id}

	// Define the update operation
	update := bson.M{"$set": bson.M{"title": newInputVideo.Title, "url": newInputVideo.URL, "author": bson.M{"name": newInputVideo.Author, "id": strconv.Itoa(rand.Int())}}}

	// Options for the FindOneAndUpdate operation
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Perform the FindOneAndUpdate operation
	var updatedDocument model.Video
	err := collection.FindOneAndUpdate(context.Background(), filter, update, options).Decode(&updatedDocument)
	if err != nil {
		log.Fatal(err)
	}

	return &updatedDocument
}
