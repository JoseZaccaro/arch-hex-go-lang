package mongo_role

import (
	"api/autentiacion/internal/domain"
	"context"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r RoleRepository) ExistsByName(name string) (bool, error) {

	filter := bson.M{"name": name}
	var result domain.CreateRole

	err := r.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
func (r RoleRepository) FindByName(name string) (*domain.Role, error) {
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: regexp.QuoteMeta(name), Options: "i"}}}
	var result domain.Role

	err := r.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
