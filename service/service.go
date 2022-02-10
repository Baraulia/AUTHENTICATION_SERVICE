package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
)

type AppUser interface {
	GetUser(id int) (*model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(id int) (*model.User, error)
	DeleteUserByID(id int) error
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
