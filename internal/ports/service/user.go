package service

import (
	"api/autentiacion/internal/domain"
	"context"
)

type UserService interface {
	Create(ctx context.Context, params domain.UserCreateParams) (*domain.User, error)
	ReadAll(ctx context.Context) ([]*domain.User, error)
	Read(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, id string, params domain.UserUpdateParams) (*domain.User, error)
	Delete(ctx context.Context, id string) error
	Register(ctx context.Context, params domain.UserCreateParams) (*domain.User, error)
	Login(ctx context.Context, username string, password string) (*domain.User, error)
}
