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

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		logger.Error("User migrate err", err)
	}

	userService := service.NewUserService(userRepository)

	userHandler := handlers.NewUserHandler(userService)

	route.HandleFunc("/user", userHandler.GetAllUser).Methods("GET")
	route.HandleFunc("/", handlers.HomePageIndex).Methods("GET")

	return route
}
