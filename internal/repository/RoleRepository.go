package repository

import (
	"github.com/stickpro/go-shop/internal/models"
	"github.com/stickpro/go-shop/pkg/logger"
	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

type RoleRepository interface {
	GetAll() ([]models.Role, error)
	Migrate() error
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{
		DB: db,
	}
}

func (r roleRepository) Migrate() error {
	logger.Info("[RoleRepository]...Migrate")
	return r.DB.AutoMigrate(&models.Role{}, &models.RoleUser{})
}

func (r roleRepository) GetAll() (role []models.Role, err error) {
	logger.Info("[RoleRepository]...Get All")
	err = r.DB.Find(&role).Error
	return role, err
}
