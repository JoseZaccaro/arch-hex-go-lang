package repository

import "api/autentiacion/internal/domain"

type RoleRepository interface {
	Insert(role *domain.CreateRole) (interface{}, error)
	// DeleteByID(id string) error
	// UpdateByID(id string, role *domain.Role) error
	// FindByID(id string) (*domain.Role, error)
	// FindAllRoles() ([]*domain.Role, error)
	FindByName(name string) (*domain.Role, error)
	ExistsByName(name string) (bool, error)
}
