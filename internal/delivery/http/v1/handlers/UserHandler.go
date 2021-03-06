package handlers

import (
	"context"
	"encoding/json"
	"github.com/stickpro/go-shop/internal/models"
	"github.com/stickpro/go-shop/internal/service"
	"github.com/stickpro/go-shop/pkg/logger"
	response "github.com/stickpro/go-shop/pkg/resource"
	"net/http"
)

type UserHandlerInterface interface {
	CreateUser(ctx context.Context)
	GetAllUser()
	TransferMoney(ctx context.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) userHandler {
	return userHandler{
		userService: s,
	}
}

func (u userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("[UserController]...add User")
	var user models.User

	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(err)
		return
	}
	logger.Info("[Request]", r.ParseForm())

	user, err := u.userService.Save(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode("Error while saving user")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u userHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("[UserController]...get all Users")

	users, err := u.userService.GetAll()
	if err != nil {
		response.Write(w,
			response.Code(http.StatusBadRequest),
			response.Error("users", "Error while saving user", map[string]string{"name":"Required field"}),
		)
		return
	}


	response.Write(w,
		response.Code(http.StatusOK),
		response.Success("users", users, nil),
	)
}
