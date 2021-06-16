package service

import (
	"github.com/stickpro/go-shop/internal/models"
	"github.com/stickpro/go-shop/internal/repository"
)

type UserService interface {
	Save(models.User) (models.User, error)
	GetAll() ([]models.User, error)
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

func (u userService) Save(user models.User) (models.User, error) {

	return u.userRepository.Save(user)
}

func (u userService) GetAll() ([]models.User, error) {

	return u.userRepository.GetAll()
}
