package service

import (
	"time"
	"user-management/model"
	"user-management/pkg/repository"
)

var usersDb = []model.User{
	{"1", "John", 23, "New-York"},
	{"2", "Sarah", 22, "Washington"},
	{"3", "Emmy", 20, "Warsaw"},
	{"4", "Sardor", 25, "Tashkent"},
}

type User interface {
	CreateUser(model.CreateUserDTO)
	GetUsers() []model.User
	GetSingleUser(string) (model.User, bool)
	UpdateUser(string, model.User)
	DeleteUser(string)
}

type UserService struct {
	userRepository repository.User
}

func NewUserService(userRepository repository.User) *UserService {
	return &UserService{
		userRepository,
	}
}

func (u *UserService) CreateUser(newUser model.CreateUserDTO) {
	newUser.Id = string(time.Now().Unix())
	usersDb = append(usersDb, model.User(newUser))
}
func (u *UserService) GetUsers() []model.User {
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
func (u *UserService) UpdateUser(id string, user model.User) {}
func (u *UserService) DeleteUser(id string)                  {}
