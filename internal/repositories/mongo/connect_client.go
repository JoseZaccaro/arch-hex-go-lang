package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(dbUrl string) (client *mongo.Client, err error) {

	// TODO: connect to mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, errCnn := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	if errCnn != nil {
		return nil, err
	}
	errCnn = client.Ping(ctx, nil)
	if errCnn != nil {
		return nil, err
	}

	return client, nil
}
