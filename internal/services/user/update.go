package user

import (
	"api/autentiacion/internal/domain"
	"context"
	"fmt"
)

func (s Service) Update(ctx context.Context, id string, params domain.UserCreateParams) (*domain.User, error) {
	if exists, _ := s.UserRepository.ExistsById(id); !exists {
		return nil, fmt.Errorf("invalid id")
	}
	role, errRole := s.RoleRepository.FindByName(params.RoleID)
	if role == nil || errRole != nil {
		return nil, fmt.Errorf("invalid role")
	}

	update := domain.UserDB{
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: params.PasswordHash,
		RoleID:       role.ID,
	}

	_, err := s.UserRepository.UpdateByID(id, &update)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user, _ := s.UserRepository.FindUserById(id)

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	userEntity := &domain.User{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		RoleID:       user.RoleID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return userEntity, nil
}
