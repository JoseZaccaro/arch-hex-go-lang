package user

import (
	"api/autentiacion/internal/domain"
	"context"
)

func (s Service) Read(ctx context.Context, id string) (*domain.UserDB, error) {
	return nil, nil
}

func (s Service) ReadAll(ctx context.Context) ([]*domain.UserDB, error) {

	users, _ := s.UserRepository.FindAllUsers()

	return users, nil
}
