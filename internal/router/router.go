package router

import (
	"github.com/gorilla/mux"
	"github.com/stickpro/go-shop/internal/delivery/http/v1/handlers"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", handlers.HomePageIndex).Methods("GET")

	return route
}
