package user

type service struct {
	repo Repository
}

//NewService create new service factory
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

// Create a User
func (s service) Create(a *User) (int, error) {
	return s.repo.Create(a)
}

// Get a User
func (s service) Get(id int) (*User, error) {
	return s.repo.Get(id)
}
