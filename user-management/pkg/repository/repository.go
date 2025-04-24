package repository

type Repository struct {
	User
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserRepository(),
	}
}
