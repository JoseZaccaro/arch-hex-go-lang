package mongo_role

import (
	"api/autentiacion/internal/domain"
	"context"
	"fmt"
	"log"
)

func (r RoleRepository) Insert(role *domain.CreateRole) (id interface{}, err error) {

	// TODO: save role
	insertResult, err2 := r.Collection.InsertOne(context.Background(), role)

	if err2 != nil {
		log.Println(err2.Error())
		return nil, fmt.Errorf("failed to insert role: %v", err2)
	}

	return insertResult.InsertedID, nil
}
