package mongo_user

import (
	"api/autentiacion/internal/domain"
	"context"
	"fmt"
	"log"
)

func (r UserRepository) Insert(user *domain.User) (id interface{}, err error) {

	// TODO: save user
	// ==========================
	// collection := r.Client.Database("mydb").Collection("users")

	insertResult, err2 := r.Collection.InsertOne(context.Background(), user)

	if err2 != nil {
		log.Println(err2.Error())
		return nil, fmt.Errorf("failed to insert user: %v", err2)
	}
	// defer r.Client.Disconnect(context.Background())

	return insertResult.InsertedID, nil
}
