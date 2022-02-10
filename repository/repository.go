package repository

import (
	"database/sql"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
)

type AppUser interface {
	GetUserByID(id int) (*model.User, error)
	GetUserAll() ([]model.User, error)
	CreateUser(User *model.User) (*model.User, error)
	UpdateUser(User model.User, id int) (*model.User, error)
	DeleteUserByID(id int) error
	GetUserByEmail(email string) (model.User, error)
}

type Repository struct {
	AppUser
}

func NewRepository(db *sql.DB, logger logging.Logger) *Repository {
	return &Repository{
		AppUser: NewUserPostgres(db, logger),
	}
}
