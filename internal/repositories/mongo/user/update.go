package mongo_user

import (
	"api/autentiacion/internal/domain"
	"api/autentiacion/internal/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r UserRepository) UpdateByID(id string, params *domain.UserDB) (interface{}, error) {
	if params == nil {
		return nil, fmt.Errorf("user create params is nil")
	}
	objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objId}}
	updates := bson.D{}
	if params.Username != "" {
		updates = append(updates, bson.E{Key: "username", Value: params.Username})
	}
	if params.Email != "" {
		updates = append(updates, bson.E{Key: "email", Value: params.Email})
	}
	if params.PasswordHash != "" {
		hashedPassword, _ := utils.HashPassword(params.PasswordHash)
		updates = append(updates, bson.E{Key: "password_hash", Value: hashedPassword})
	}
	if params.RoleID != "" {
		updates = append(updates, bson.E{Key: "role_id", Value: params.RoleID})
	}

	now := time.Now().UTC()

	updates = append(updates, bson.E{Key: "updated_at", Value: now})

	updateDocument := bson.D{
		{Key: "$set", Value: updates},
	}
	ctx := context.Background()
	// user := &domain.UserMongo{}
	result, err := r.Collection.UpdateOne(ctx, filter, updateDocument)
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}
