package repository

import (
	"database/sql"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go

type AppUser interface {
	GetUserByID(id int) (*model.ResponseUser, error)
	GetUserAll(page int, limit int) ([]model.ResponseUser, int, error)
	CreateUser(User *model.CreateUser) (int, error)
	UpdateUser(User *model.UpdateUser, id int) error
	DeleteUserByID(id int) (int, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserPasswordByID(id int) (string, error)
}

type Repository struct {
	AppUser
}

func NewRepository(db *sql.DB, logger logging.Logger) *Repository {
	return &Repository{
		AppUser: NewUserPostgres(db, logger),
	}
}
