package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", HomeHandler)

	return route
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Test: %v\n", vars["Test"])
}
