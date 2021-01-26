package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/chatApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPeer(phNum string, db *mongo.Database) (models.Peer, error) {
	fmt.Println(phNum)
	filter := bson.D{{"ph", phNum}}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results models.Peer
	collection := db.Collection("peers")
	err := collection.FindOne(ctx, filter).Decode(&results)
	if err != nil {
		return models.Peer{}, nil
	}
	return results, nil
}
