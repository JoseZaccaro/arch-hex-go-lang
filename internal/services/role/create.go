package role

import (
	"api/autentiacion/internal/domain"
	"context"
	"fmt"
	"time"
)

func (s Service) Create(ctx context.Context, params domain.CreateRole) (interface{}, error) {
	now := time.Now().UTC()

	if params.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if exists, err := s.RoleRepository.ExistsByName(params.Name); exists {
		return nil, fmt.Errorf("role already exists")
	} else if err != nil {
		return nil, err
	}

	role := &domain.CreateRole{
		Name:        params.Name,
		Description: params.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	id, errInsrt := s.RoleRepository.Insert(role)

	if errInsrt != nil {
		fmt.Println(errInsrt)
		return nil, errInsrt
	}

	return id, nil

}
