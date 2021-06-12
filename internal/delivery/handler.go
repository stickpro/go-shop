package delivery

import "github.com/stickpro/go-shop/internal/service"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}
