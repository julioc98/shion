package user

// Repository User interface
type Repository interface {
	Create(user *User) (int, error)
	Get(id int) (*User, error)
}

// Service User interface
type Service interface {
	Create(user *User) (int, error)
	Get(id int) (*User, error)
}
