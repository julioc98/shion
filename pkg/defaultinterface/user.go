package defaultinterface

import (
	"net/http"

	"github.com/julioc98/shion/pkg/entity"
)

// UserRepository represent the User's repository default contract
type UserRepository interface {
	Create(e *entity.User) (*entity.User, error)
	GetByID(id uint) (*entity.User, error)
	Update(e *entity.User) (*entity.User, error)
	Delete(e *entity.User) error
	FindOne(query *entity.User, args ...string) (*entity.User, error)
	FindMany(query *entity.User, args ...string) ([]entity.User, error)
	FindAll() ([]entity.User, error)
}

// UserUseCase represent the User's Use Case default contract
type UserUseCase interface {
	Create(e *entity.User) (*entity.User, error)
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetByEmailAndPassword(email, password string) (*entity.User, error)
}

// UserHTTPHandler represent the User's http handler default Interface
type UserHTTPHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
}
