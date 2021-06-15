package repository

import (
	"errors"
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
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
	WithTrx(*gorm.DB) userRepository
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

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
	if trxHandle == nil {
		logger.Info("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}

func (u userRepository) IncrementMoney(receiver uint, amount float64) error {
	logger.Info("[UserRepository]...Increment Money")
	return u.DB.Model(&models.User{}).Where("id=?", receiver).Update("wallet", gorm.Expr("wallet + ?", amount)).Error
}

func (u userRepository) DecrementMoney(giver uint, amount float64) error {
	logger.Info("[UserRepository]...Decrement Money")
	return errors.New("something")
	// return u.DB.Model(&model.User{}).Where("id=?", giver).Update("wallet", gorm.Expr("wallet - ?", amount)).Error
}
