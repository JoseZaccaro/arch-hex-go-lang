package mongo_user

import (
	"api/autentiacion/internal/domain"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r UserRepository) FindUserByUsername(username string) (*domain.UserMongo, error) {
	var result bson.M

	// TODO: find user by username
	err := r.Collection.FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found")
		}
		return nil, err
	}

	user := &domain.UserMongo{
		ID:           result["_id"].(primitive.ObjectID),
		Username:     result["username"].(string),
		Email:        result["email"].(string),
		PasswordHash: result["password_hash"].(string),
		CreatedAt:    result["created_at"].(primitive.DateTime).Time(),
		UpdatedAt:    result["updated_at"].(primitive.DateTime).Time(),
	}
	return user, nil
}
func (r UserRepository) FindUserByEmail(email string) (*domain.UserMongo, error) {
	var result bson.M

	// TODO: find user by email
	err := r.Collection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found")
		}
		return nil, err
	}

	user := &domain.UserMongo{
		ID:           result["_id"].(primitive.ObjectID),
		Username:     result["username"].(string),
		Email:        result["email"].(string),
		PasswordHash: result["password_hash"].(string),
		CreatedAt:    result["created_at"].(primitive.DateTime).Time(),
		UpdatedAt:    result["updated_at"].(primitive.DateTime).Time(),
	}
	return user, nil
}

func (r UserRepository) ExistsByEmail(email string) (bool, error) {

	var result bson.M
	err := r.Collection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r UserRepository) ExistsByUsername(username string) (bool, error) {

	var result bson.M
	err := r.Collection.FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
