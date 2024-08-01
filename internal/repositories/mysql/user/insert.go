package mysql_user

import (
	"api/autentiacion/internal/domain"
	"api/autentiacion/internal/repositories"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r UserRepository) Insert(user *domain.User) (id interface{}, err error) {
	mapper := &repositories.Mapper{}
	now := time.Now()

	roleID := user.RoleID.(primitive.ObjectID)
	roleStr := roleID.Hex()
	user.RoleID = roleStr

	newUser := &domain.UserDB{
		ID:           nil,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		RoleID:       user.RoleID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	ctx := r.DB.Table("users").Create(newUser)
	errIns := ctx.Error

	if errIns != nil {
		log.Println(errIns)
		return nil, fmt.Errorf("failed to insert user: %v", errIns)
	}

	rows, _ := r.DB.Table("users").Select("id").Order("created_at DESC").Limit(1).Rows()

	if rows.Next() {
		rows.Scan(&id)
	}

	return mapper.ToInt64(id), nil

}
