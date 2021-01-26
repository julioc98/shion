package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// User entity
type User struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Name      string         `gorm:"not null;" json:"name" validate:"required"`
	Email     string         `gorm:"index:idx_email_password;index:idx_email;unique;not null;" json:"email" validate:"required,email"`
	Password  string         `gorm:"index:idx_email_password;not null;" json:"password,omitempty" validate:"required"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Active    sql.NullBool   `gorm:"default:true" json:"active,omitempty"`
}

// OmitPassword remove password
func (u *User) OmitPassword() {
	u.Password = ""
}
