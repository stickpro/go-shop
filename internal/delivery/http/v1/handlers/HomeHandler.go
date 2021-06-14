package handlers

import (
	"encoding/json"
	"net/http"
)

func HomePageIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
		Age  int
	}{"Joh Doe", 30}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
