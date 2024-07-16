package repository

import "api/autentiacion/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	UpdateByID(id string, user *domain.User) error
	DeleteByID(id string) error
	FindAllUsers() ([]*domain.User, error)
	FindUserByUsername(username string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
}
