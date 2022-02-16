package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
)

//go:generate mockgen -source = service.go -destination = mocks/service_mock.go

type AppUser interface {
	GetUser(id int) (*model.User, error)
	GetUsers(page int, limit int) ([]model.User, error)
	CreateUser(user *model.CreateUser) (*model.User, error)
	UpdateUser(user *model.UpdateUser, id int) (int, error)
	DeleteUserByID(id int) (int, error)
	AuthUser(email string, password string) (int, error)
}

type Service struct {
	AppUser
}

func NewService(rep *repository.Repository, logger logging.Logger) *Service {
	return &Service{
		AppUser: NewUserService(*rep, logger),
	}
}
