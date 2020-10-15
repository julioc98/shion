package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julioc98/shion/internal/app/user"
)

// UserHandler Interface
type UserHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService user.Service
}

// NewUserHandler Create a new handler
func NewUserHandler(userService user.Service) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// Add a User
func (ah *userHandler) Add(w http.ResponseWriter, r *http.Request) {
	var req user.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := ah.userService.Create(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "user_id": %d }`, id)))
}

// FindByID a User
func (ah *userHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user, err := ah.userService.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
