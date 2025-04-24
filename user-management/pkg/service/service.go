package service

import "user-management/pkg/repository"

type Service struct {
	User
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repositories.User),
	}
}
