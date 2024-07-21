package service

import (
	"api/autentiacion/internal/domain"
	"context"
)

type RoleService interface {
	Create(ctx context.Context, params domain.CreateRole) (interface{}, error)
	ReadAll(ctx context.Context) ([]*domain.Role, error)
	FindByName(ctx context.Context, name string) (*domain.Role, error)
	Update(ctx context.Context, id string, params domain.CreateRole) (*domain.Role, error)
	Delete(ctx context.Context, id string) error
}
