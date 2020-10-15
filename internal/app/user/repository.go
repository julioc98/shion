package user

import (
	gorm "github.com/jinzhu/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

//NewPostgresRepository create new postgres repository
func NewPostgresRepository(db *gorm.DB) Repository {
	return &postgresRepository{
		db,
	}
}

// Create User
func (r *postgresRepository) Create(a *User) (int, error) {
	if dbc := r.db.Create(a); dbc.Error != nil {
		return 0, dbc.Error
	}
	return a.ID, nil
}

// Get User
func (r *postgresRepository) Get(id int) (*User, error) {
	var user User
	if dbc := r.db.First(&user, id); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &user, nil
}
