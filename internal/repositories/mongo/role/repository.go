package mongo_role

import "go.mongodb.org/mongo-driver/mongo"

type RoleRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}
