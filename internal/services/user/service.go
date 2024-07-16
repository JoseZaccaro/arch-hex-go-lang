package user

import (
	"api/autentiacion/internal/ports/repository"
	"api/autentiacion/internal/ports/service"
)

// Service implements service.UserService

type Service struct {
	UserRepository repository.UserRepository
}

func NewService() service.UserService {
	return &Service{}
}
