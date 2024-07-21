package user

import (
	"api/autentiacion/internal/domain"
	"api/autentiacion/internal/utils"
	"context"
	"fmt"

	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s Service) Create(ctx context.Context, params domain.UserCreateParams) (id interface{}, err error) {
	now := time.Now().UTC()
	hashedPassword, _ := utils.HashPassword(params.PasswordHash)

	if params.RoleID == "" {
		params.RoleID = "demo"
	}

	user := &domain.User{
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: string(hashedPassword),
		RoleID:       params.RoleID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	id, errInsrt := s.UserRepository.Insert(user)

	if errInsrt != nil {
		fmt.Println(errInsrt)
		return nil, errInsrt
	}

	return id, nil
}

func (s Service) Register(ctx context.Context, params domain.UserCreateParams) (*domain.UserMongo, error) {
	now := time.Now().UTC()
	exists, err := s.UserRepository.ExistsByEmail(params.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPassword, _ := utils.HashPassword(params.PasswordHash)
	role, _ := s.RoleRepository.FindByName("user")

	id, errInsrt := s.UserRepository.Insert(&domain.User{
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: string(hashedPassword),
		RoleID:       role.ID,
		CreatedAt:    now,
		UpdatedAt:    now,
	})

	if errInsrt != nil {
		return nil, err
	}

	newUser := &domain.UserMongo{
		ID:           id,
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: params.PasswordHash,
		RoleID:       params.RoleID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return newUser, nil
}

func (s Service) Login(ctx context.Context, email string, password string) (*domain.UserLogin, error) {
	exists, err := s.UserRepository.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("email or password incorrect")
	}

	user, err := s.UserRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, fmt.Errorf("email or password incorrect")
	}

	userLogin := &domain.UserLogin{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return userLogin, nil
}
