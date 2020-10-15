package router

import (
	"github.com/gorilla/mux"
	"github.com/julioc98/shion/cmd/api/handler"
)

// SetUserRoutes add routes from User
func SetUserRoutes(ah handler.UserHandler, r *mux.Router) {
	r.HandleFunc("", ah.Add).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", ah.FindByID).Methods("GET")
}
