package mongo_user

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
