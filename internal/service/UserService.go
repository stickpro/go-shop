package service

import (
	"github.com/stickpro/go-shop/internal/models"
	"github.com/stickpro/go-shop/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	Save(models.User) (models.User, error)
	GetAll() ([]models.User, error)
	WithTrx(*gorm.DB) userService
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService -> returns new user service
func NewUserService(r repository.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

// WithTrx enables repository with transaction
func (u userService) WithTrx(trxHandle *gorm.DB) userService {
	u.userRepository = u.userRepository.WithTrx(trxHandle)
	return u
}

func (u userService) Save(user models.User) (models.User, error) {

	return u.userRepository.Save(user)
}

func (u userService) GetAll() ([]models.User, error) {

	return u.userRepository.GetAll()
}

func (u userService) IncrementMoney(receiver uint, amount float64) error {

	return u.userRepository.IncrementMoney(receiver, amount)
}

func (u userService) DecrementMoney(giver uint, amount float64) error {

	return u.userRepository.DecrementMoney(giver, amount)
}
