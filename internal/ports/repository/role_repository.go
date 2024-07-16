package repository

import "api/autentiacion/internal/domain"

type RoleRepository interface {
	Save(role *domain.Role) error
	DeleteByID(id string) error
	UpdateByID(id string, role *domain.Role) error
	FindAllRoles() ([]*domain.Role, error)
	FindRoleByName(name string) (*domain.Role, error)
}
