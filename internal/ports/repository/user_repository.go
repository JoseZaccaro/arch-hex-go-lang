package repository

import "api/autentiacion/internal/domain"

type UserRepository interface {
	Insert(user *domain.User) (id interface{}, err error)
	UpdateByID(id string, user *domain.UserDB) (interface{}, error)
	// DeleteByID(id string) error
	FindAllUsers() ([]*domain.UserDB, error)
	FindUserByUsername(username string) (*domain.UserDB, error)
	FindUserByEmail(email string) (*domain.UserDB, error)
	FindUserById(id string) (*domain.UserDB, error)
	ExistsByEmail(email string) (bool, error)
	ExistsById(id string) (bool, error)
	ExistsByUsername(username string) (bool, error)
}
