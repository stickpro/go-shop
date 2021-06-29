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
	Migrate() error
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{
		DB: db,
	}
}

func (r roleRepository) Migrate() error {
	logger.Info("[RoleRepository]...Migrate")

	return r.DB.AutoMigrate(&models.Role{})
}
func (r roleRepository) MakePivotRoleUser() error {
	logger.Info("[RoleRepository]...Migrate Pivot User Role")
	err := r.DB.AutoMigrate(&models.RoleUser{})
	if err != nil {
		logger.Error("[RoleRepository]... Pivot Error Migrate")
	}
}
