package repository

import (
	"github.com/stickpro/go-shop/internal/models"
	"github.com/stickpro/go-shop/pkg/logger"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Save(models.User) (models.User, error)
	GetAll() ([]models.User, error)
	Migrate() error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}
func (u userRepository) Migrate() error {
	logger.Info("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&models.User{})
}

func (u userRepository) Save(user models.User) (models.User, error) {
	logger.Info("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) GetAll() (users []models.User, err error) {
	logger.Info("[UserRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err
}
