package service

import (
	"github.com/stickpro/go-shop/internal/models"
	"github.com/stickpro/go-shop/internal/repository"
)

type RoleService interface {
	GetAll() ([]models.Role, error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(r repository.RoleRepository) RoleService {
	return roleService{
		roleRepository: r,
	}
}

func (r roleService) GetAll() ([]models.Role, error) {
	return r.roleRepository.GetAll()
}
