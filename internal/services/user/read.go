package user

import (
	"api/autentiacion/internal/domain"
	"context"
)

func (s Service) Read(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (s Service) ReadAll(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}
