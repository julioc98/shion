package repository

import (
	"github.com/julioc98/shion/pkg/entity"
	"gorm.io/gorm"
)

// User repository
type User struct {
	db *gorm.DB
}

//NewUserRepository create new postgres repository
func NewUserRepository(db *gorm.DB) *User {
	return &User{
		db,
	}
}

// Create User
func (r *User) Create(e *entity.User) (*entity.User, error) {
	if dbc := r.db.Create(e); dbc.Error != nil {
		return nil, dbc.Error
	}
	return e, nil
}

// GetByID User
func (r *User) GetByID(id uint) (*entity.User, error) {
	e := &entity.User{}
	if dbc := r.db.First(e, id); dbc.Error != nil {
		return nil, dbc.Error
	}
	return e, nil
}

// Update User
func (r *User) Update(e *entity.User) (*entity.User, error) {
	if dbc := r.db.Save(e); dbc.Error != nil {
		return nil, dbc.Error
	}
	return e, nil
}

// Delete User
func (r *User) Delete(e *entity.User) error {
	if dbc := r.db.Delete(e); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

// FindOne User
func (r *User) FindOne(query *entity.User, args ...string) (*entity.User, error) {
	e := &entity.User{}
	if dbc := r.db.Where(query, args).First(e); dbc.Error != nil {
		return nil, dbc.Error
	}

	return e, nil
}

// FindMany Users
func (r *User) FindMany(query *entity.User, args ...string) ([]entity.User, error) {
	var e []entity.User
	if dbc := r.db.Where(query, args).Find(e); dbc.Error != nil {
		return nil, dbc.Error
	}
	return e, nil
}

// FindAll Users
func (r *User) FindAll() ([]entity.User, error) {
	var e []entity.User
	if dbc := r.db.Find(e); dbc.Error != nil {
		return nil, dbc.Error
	}
	return e, nil
}
