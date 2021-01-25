package main

import (
	"context"
	"log"
	"time"

	"github.com/chatApp/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var s server.Server
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://anmol:anmol@cluster0.cnhws.mongodb.net/"))
	HandleError(err)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	HandleError(err)
	db := client.Database("falconet")
	s = server.New(db)
	defer client.Disconnect(ctx)
	s.Listen(":8080")
}
