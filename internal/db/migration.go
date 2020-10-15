package db

import (
	"github.com/jinzhu/gorm"
	"github.com/julioc98/shion/internal/app/user"
)

// Migrate migration BD
func Migrate(conn *gorm.DB) {
	// Migrate the schema
	conn.AutoMigrate(&user.User{})

	// Create an User
	conn.Create(&user.User{Email: "init@shion.com", Password: "123"})

}
