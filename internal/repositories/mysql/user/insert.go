package mysql_user

import (
	"api/autentiacion/internal/domain"
	"fmt"
	"log"
)

func (r UserRepository) Insert(user *domain.User) (id interface{}, err error) {
	// TODO: save user
	ctx := r.DB.Table("users").Create(user)
	errIns := ctx.Error

	if errIns != nil {
		log.Println(errIns.Error())
		return nil, fmt.Errorf("failed to insert user: %v", errIns)
	}
	// defer r.Client.Disconnect(context.Background())

	insertResult := ctx.RowsAffected

	return insertResult, nil

}
