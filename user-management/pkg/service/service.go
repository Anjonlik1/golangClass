package service

type Service struct {
	User
}

func NewService() *Service {
	return &Service{
		User: NewUserService(),
	}

}
