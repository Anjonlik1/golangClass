package repository

import "user-management/model"

type User interface {
	InsertUser(model.User) error
	FetchUsers() ([]model.User, error)
	FetchSingleUser() (model.User, error)
	UpdateUser(model.User) error
	DeleteUser(string) error
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) InsertUser(newUser model.User) error {
	return nil
}

func (u *UserRepository) FetchUsers() ([]model.User, error) {
	return []model.User{}, nil
}

func (u *UserRepository) FetchSingleUser() (model.User, error) {
	return model.User{}, nil
}

func (u *UserRepository) UpdateUser(user model.User) error {
	return nil
}

func (u *UserRepository) DeleteUser(id string) error {
	return nil
}
