package service

import (
	"api/autentiacion/internal/domain"
	"context"
)

type UserService interface {
	Create(ctx context.Context, params domain.UserCreateParams) (interface{}, error)
	ReadAll(ctx context.Context) ([]*domain.UserDB, error)
	Read(ctx context.Context, id string) (*domain.UserDB, error)
	Update(ctx context.Context, id string, params domain.UserCreateParams) (*domain.User, error)
	Delete(ctx context.Context, id string) error
	Register(ctx context.Context, params domain.UserCreateParams) (*domain.UserDB, error)
	Login(ctx context.Context, username string, password string) (*domain.UserLogin, error)
}
