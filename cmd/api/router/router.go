package router

import (
	"github.com/gorilla/mux"
	"github.com/julioc98/shion/pkg/defaultinterface"
)

// SetUserRoutes add routes from User
func SetUserRoutes(ah defaultinterface.UserHTTPHandler, r *mux.Router) {
	r.HandleFunc("", ah.Create).Methods("POST")
	r.HandleFunc("", ah.Get).Methods("GET")
	r.HandleFunc("/auth", ah.GetToken).Methods("GET")
}
