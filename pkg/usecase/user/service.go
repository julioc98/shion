package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julioc98/shion/pkg/defaultinterface"
	"github.com/julioc98/shion/pkg/entity"
	"github.com/shaj13/go-guardian/v2/auth"
)

// Service User usecase
type Service struct {
	repo        defaultinterface.UserRepository
	passService defaultinterface.PasswordService
}

//NewService create new service (factory)
func NewService(r defaultinterface.UserRepository, p defaultinterface.PasswordService) *Service {
	return &Service{
		repo:        r,
		passService: p,
	}
}

// Create a new User
func (s *Service) Create(e *entity.User) (*entity.User, error) {
	pass, err := s.passService.Generate(e.Password)
	if err != nil {
		return nil, err
	}
	e.Password = pass
	return s.repo.Create(e)
}

// GetByID  Get Users By ID
func (s *Service) GetByID(id uint) (*entity.User, error) {
	return s.repo.GetByID(id)
}

// GetByEmail Get User By Email
func (s *Service) GetByEmail(email string) (*entity.User, error) {
	e := &entity.User{
		Email: email,
	}
	return s.repo.FindOne(e, "email")
}

// GetByEmailAndPassword  Get User By Email And Password
func (s *Service) GetByEmailAndPassword(email, password string) (*entity.User, error) {
	e := &entity.User{
		Email:    email,
		Password: password,
	}
	pass, err := s.passService.Generate(e.Password)
	if err != nil {
		return nil, err
	}
	e.Password = pass
	return s.repo.FindOne(e, "email", "password")
}

// ValidateEmailAndPassword Compare Email and Password
func (s *Service) ValidateEmailAndPassword(ctx context.Context, r *http.Request, email, password string) (auth.Info, error) {
	e, err := s.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	pass, err := s.passService.Generate(e.Password)
	if err != nil {
		return nil, err
	}

	err = s.passService.Compare(pass, e.Password)
	if err != nil {
		return nil, err
	}

	strID := fmt.Sprintf("%d", e.ID)
	return auth.NewDefaultUser(e.Name, strID, nil, nil), nil
}
