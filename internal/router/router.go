package router

import (
	"github.com/gorilla/mux"
	"github.com/stickpro/go-shop/internal/delivery/http/v1/handlers"
	"github.com/stickpro/go-shop/internal/repository"
	"github.com/stickpro/go-shop/internal/service"
	"github.com/stickpro/go-shop/pkg/logger"
	"gorm.io/gorm"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init(db *gorm.DB) *mux.Router {
	route := mux.NewRouter()

	userRepository := InitUserRepository(db)
	//roleRepository := InitRoleRepository(db)

	//roleService := service.NewRoleService(roleRepository)

	userService := service.NewUserService(userRepository)

	userHandler := handlers.NewUserHandler(userService)

	route.HandleFunc("/user", userHandler.GetAllUser).Methods("GET")
	route.HandleFunc("/user", userHandler.AddUser).Methods("POST")
	route.HandleFunc("/", handlers.HomePageIndex).Methods("GET")

	return route
}

func InitUserRepository(db *gorm.DB) repository.UserRepository {
	userRepository := repository.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		logger.Error("User migrate err", err)
	}
	return userRepository
}

func InitRoleRepository(db *gorm.DB) repository.RoleRepository {
	roleRepository := repository.NewRoleRepository(db)
	if err := roleRepository.Migrate(); err != nil {
		logger.Error("Role migrate err", err)
	}
	return roleRepository
}
