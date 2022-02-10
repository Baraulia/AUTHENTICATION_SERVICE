package repository

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserPostgres struct {
	db     *sql.DB
	logger logging.Logger
}

func NewUserPostgres(db *sql.DB, logger logging.Logger) *UserPostgres {
	return &UserPostgres{db: db, logger: logger}
}

// GetUserByID ...
func (u UserPostgres) GetUserByID(id int) (*model.User, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("GetUserByID: can not starts transaction:%s", err)
		return nil, fmt.Errorf("getUserByID: can not starts transaction:%w", err)
	}
	var user model.User
	result := transaction.QueryRow("SELECT id, email, password, created_at FROM users WHERE id = $1", id)
	if err := result.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		u.logger.Errorf("GetUserByID: error while scanning for user:%s", err)
		return nil, fmt.Errorf("getUserByID: repository error:%w", err)
	}

	return &user, transaction.Commit()
}

// GetUserAll ...
func (u *UserPostgres) GetUserAll(page int, limit int) ([]model.User, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("GetUserAll: can not starts transaction:%s", err)
		return nil, fmt.Errorf("getUserAll: can not starts transaction:%w", err)
	}
	var User model.User
	var Users []model.User
	var query string
	if page == 0 || limit == 0 {
		query = "SELECT id, email, password, created_at FROM users"
	} else {
		query = fmt.Sprintf("SELECT id, email, password, created_at FROM users LIMIT %d OFFSET %d", limit, (page-1)*limit)
	}

	rows, err := transaction.Query(query)
	if err != nil {
		u.logger.Errorf("GetUserAll: can not executes a query:%s", err)
		return nil, fmt.Errorf("getUserAll:repository error:%w", err)
	}

	for rows.Next() {
		if err := rows.Scan(&User.ID, &User.Email, &User.Password, &User.CreatedAt); err != nil {
			u.logger.Errorf("Error while scanning for user:%s", err)
			return nil, fmt.Errorf("getUserAll:repository error:%w", err)
		}
		Users = append(Users, User)
	}
	return Users, transaction.Commit()
}

// CreateUser ...
func (u *UserPostgres) CreateUser(user *model.CreateUser) (*model.User, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("CreateUser: can not starts transaction:%s", err)
		return nil, fmt.Errorf("createUser: can not starts transaction:%w", err)
	}
	var createdUser model.User
	defer transaction.Rollback()
	hash, err := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Errorf("CreateUser: can not generate hash from password:%s", err)
		return nil, fmt.Errorf("createUser: can not generate hash from password:%w", err)
	}
	row := transaction.QueryRow("INSERT INTO users (email, password, created_at) VALUES ($1, $2, $3) RETURNING id, email, password, created_at", user.Email, hash, time.Now())
	if err := row.Scan(&createdUser.ID, &createdUser.Email, &createdUser.Password, &createdUser.CreatedAt); err != nil {
		u.logger.Errorf("CreateUser: error while scanning for user:%s", err)
		return nil, fmt.Errorf("createUser: error while scanning for user:%w", err)
	}
	return &createdUser, transaction.Commit()
}

// UpdateUser ...
func (u *UserPostgres) UpdateUser(user *model.UpdateUser, id int) (int, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("UpdateUser: can not starts transaction:%s", err)
		return 0, fmt.Errorf("updateUser: can not starts transaction:%w", err)
	}
	defer transaction.Rollback()
	var userId int
	row := transaction.QueryRow("UPDATE users SET password =$1 WHERE id=$2 RETURNING id", user.NewPassword, id)
	if err := row.Scan(&userId); err != nil {
		u.logger.Errorf("UpdateUser: error while scanning for user:%s", err)
		return 0, fmt.Errorf("updateUser: error while scanning for user:%w", err)
	}
	return userId, transaction.Commit()
}

// DeleteUserByID ...
func (u *UserPostgres) DeleteUserByID(id int) (int, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("DeleteUserByID: can not starts transaction:%s", err)
		return 0, fmt.Errorf("deleteUserByID: can not starts transaction:%w", err)
	}
	defer transaction.Rollback()
	var userId int
	row := transaction.QueryRow("DELETE FROM users WHERE id=$1 RETURNING id", id)
	if err := row.Scan(&userId); err != nil {
		u.logger.Errorf("DeleteUserByID: error while scanning for userId:%s", err)
		return 0, fmt.Errorf("deleteUserByID: error while scanning for userId:%w", err)
	}
	return userId, transaction.Commit()
}

// GetUserByEmail ...
func (u *UserPostgres) GetUserByEmail(email string) (*model.User, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("GetUserByEmail: can not starts transaction:%s", err)
		return nil, fmt.Errorf("getUserByEmail: can not starts transaction:%w", err)
	}
	var User model.User
	query := "SELECT id, email, password, created_at FROM users WHERE email = $1"
	row := transaction.QueryRow(query, email)
	if err := row.Scan(&User.ID, &User.Email, &User.Password, &User.CreatedAt); err != nil {
		u.logger.Errorf("Error while scanning for user:%s", err)
		return nil, fmt.Errorf("getUserByEmail: repository error:%w", err)

	}
	return &User, transaction.Commit()
}
