package role

import (
	"api/autentiacion/internal/ports/repository"
	"api/autentiacion/internal/ports/service"
)

type Service struct {
	RoleRepository repository.RoleRepository
}

func NewService() service.RoleService {
	return &Service{}
}
