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

func (r UserRepository) FindUserByUsername(username string) (*domain.UserDB, error) {
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

	user := &domain.UserDB{
		ID:           result["_id"].(primitive.ObjectID),
		Username:     result["username"].(string),
		RoleID:       result["role_id"].(primitive.ObjectID),
		Email:        result["email"].(string),
		PasswordHash: result["password_hash"].(string),
		CreatedAt:    result["created_at"].(primitive.DateTime).Time(),
		UpdatedAt:    result["updated_at"].(primitive.DateTime).Time(),
	}
	return user, nil
}
func (r UserRepository) FindUserByEmail(email string) (*domain.UserDB, error) {
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

	user := &domain.UserDB{
		ID:           result["_id"].(primitive.ObjectID),
		Username:     result["username"].(string),
		Email:        result["email"].(string),
		RoleID:       result["role_id"].(primitive.ObjectID),
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
func (r UserRepository) ExistsById(id string) (bool, error) {

	var result bson.M
	objId, _ := primitive.ObjectIDFromHex(id)
	err := r.Collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: objId}}).Decode(&result)
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
func (r UserRepository) FindUserById(id string) (*domain.UserDB, error) {
	var userResult bson.M
	objId, _ := primitive.ObjectIDFromHex(id)

	err := r.Collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: objId}}).Decode(&userResult)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, err
	}

	user := &domain.UserDB{
		ID:           userResult["_id"].(primitive.ObjectID),
		Username:     userResult["username"].(string),
		Email:        userResult["email"].(string),
		RoleID:       userResult["role_id"].(primitive.ObjectID),
		PasswordHash: userResult["password_hash"].(string),
		CreatedAt:    userResult["created_at"].(primitive.DateTime).Time(),
		UpdatedAt:    userResult["updated_at"].(primitive.DateTime).Time(),
	}

	return user, nil
}

func (r UserRepository) FindAllUsers() ([]*domain.UserDB, error) {
	var users []*domain.UserDB
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var user bson.M
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &domain.UserDB{
			ID:           user["_id"].(primitive.ObjectID),
			Username:     user["username"].(string),
			RoleID:       user["role_id"].(primitive.ObjectID),
			Email:        user["email"].(string),
			PasswordHash: user["password_hash"].(string),
			CreatedAt:    user["created_at"].(primitive.DateTime).Time(),
			UpdatedAt:    user["updated_at"].(primitive.DateTime).Time(),
		})
	}
	return users, nil
}
