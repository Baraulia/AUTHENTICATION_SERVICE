package repository

import (
	"database/sql"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
)

type AppUser interface {
	GetUserByID(id int) (*model.User, error)
	GetUserAll(page int, limit int) ([]model.User, error)
	CreateUser(User *model.CreateUser) (*model.User, error)
	UpdateUser(User *model.UpdateUser, id int) (int, error)
	DeleteUserByID(id int) (int, error)
	GetUserByEmail(email string) (*model.User, error)
}

type Repository struct {
	AppUser
}

func NewRepository(db *sql.DB, logger logging.Logger) *Repository {
	return &Repository{
		AppUser: NewUserPostgres(db, logger),
	}
}
