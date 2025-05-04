package repository

import "user-management/model"

type User interface {
	InsertUser(model.User) error
	FetchUsers() ([]model.User, error)
	FetchSingleUser(id string) (model.User, error)
	UpdateUser(id string, user model.User) error
	DeleteUser(id string) error
}
type UserRepository struct {
	User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) InsertUser(model.User) error {
	return nil
}
func (u *UserRepository) FetchUsers() ([]model.User, error) {
	return nil, nil
}
func (u *UserRepository) FetchSingleUser(id string) (model.User, error) {
	return model.User{}, nil
}
func (u *UserRepository) UpdateUser(id string, user model.User) error {
	return nil
}
func (u *UserRepository) DeleteUser(id string) error {
	return nil
}
