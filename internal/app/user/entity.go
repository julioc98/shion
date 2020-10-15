package user

// User entity
type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
