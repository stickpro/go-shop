package router

import (
	"github.com/gorilla/mux"
	"github.com/stickpro/go-shop/internal/delivery/http/v1/controllers"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", controllers.HomePageIndex)

	return route
}
