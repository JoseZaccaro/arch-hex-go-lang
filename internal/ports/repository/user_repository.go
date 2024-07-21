package repository

import "api/autentiacion/internal/domain"

type UserRepository interface {
	Insert(user *domain.User) (id interface{}, err error)
	UpdateByID(id string, user *domain.UserMongo) (interface{}, error)
	// DeleteByID(id string) error
	FindAllUsers() ([]*domain.UserMongo, error)
	FindUserByUsername(username string) (*domain.UserMongo, error)
	FindUserByEmail(email string) (*domain.UserMongo, error)
	FindUserById(id string) (*domain.UserMongo, error)
	ExistsByEmail(email string) (bool, error)
	ExistsById(id string) (bool, error)
	ExistsByUsername(username string) (bool, error)
}
