package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(dbUrl string) (client *mongo.Client, err error) {

	// TODO: connect to mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dbUrl)
	clientDB, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = clientDB.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return clientDB, nil
}
