package role

import (
	"api/autentiacion/internal/domain"
	"context"
)

func (s Service) ReadAll(ctx context.Context) ([]*domain.Role, error) {
	roles, _ := s.RoleRepository.FindAllRoles()

	return roles, nil
}
func (s Service) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	return nil, nil
}
