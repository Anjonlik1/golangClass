package service

import (
	"fmt"
	"time"
	"user-management/model"
)

var usersDb = []model.User{
	{
		Id:      "1",
		Name:    "Tuhtaboy",
		Age:     30,
		Address: "123 Main St"},
	{
		Id:      "2",
		Name:    "Ali",
		Age:     25,
		Address: "456 Elm St"},
	{
		Id:      "3",
		Name:    "Bob",
		Age:     35,
		Address: "789 Oak St"},
	{
		Id:      "4",
		Name:    "Alice",
		Age:     28,
		Address: "101 Pine St"},
}

type User interface {
	CreateUser(model.CreateUserDTO)
	GetUser() []model.User
	GetSingleUser(id string) (model.User, bool)
	UpdateUser(id string, user model.User)
	DeleteUser(id string) error
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(newUser model.CreateUserDTO) {
	newUser.Id = string(time.Now().Unix())
	usersDb = append(usersDb, model.User(newUser))

}
func (u *UserService) GetUser() []model.User {
	return usersDb

}
func (u *UserService) GetSingleUser(id string) (model.User, bool) {

	for _, u := range usersDb {
		if u.Id == id {
			return u, true
		}
	}
	return model.User{}, false

}

func (u *UserService) UpdateUser(id string, user model.User) {
	for i, u := range usersDb {
		if u.Id == id {
			usersDb[i] = user
			return
		}
	}
	// If user not found, you might want to handle this case
}
func (u *UserService) DeleteUser(id string) error {
	for i, u := range usersDb {
		if u.Id == id {
			usersDb = append(usersDb[:i], usersDb[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id %s not found", id)
}
