package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/julioc98/shion/pkg/defaultinterface"
	"github.com/julioc98/shion/pkg/entity"
	"github.com/julioc98/shion/pkg/guardian"
)

// UserHTTPHandler ...
type UserHTTPHandler struct {
	userService defaultinterface.UserUseCase
	guard       *guardian.Guardian
	validate    *validator.Validate
}

// NewUserHTTPHandler Create a new handler
func NewUserHTTPHandler(userService defaultinterface.UserUseCase, validate *validator.Validate, guard *guardian.Guardian) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
		validate:    validate,
		guard:       guard,
	}
}

// Create an User
func (ah *UserHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ah.validate.Struct(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := ah.userService.Create(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.OmitPassword()

	res, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

//GetToken generate JWT Token
func (ah *UserHTTPHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ah.guard.CreateToken(w, r)
}

//Get an User
func (ah *UserHTTPHandler) Get(w http.ResponseWriter, r *http.Request) {

	idStr := ah.guard.GetUserID(r)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, err := ah.userService.GetByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user.OmitPassword()

	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

//GetByID a User
func (ah *UserHTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, err := ah.userService.GetByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user.OmitPassword()

	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
