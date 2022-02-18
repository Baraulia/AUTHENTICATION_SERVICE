package repository

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
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
func (u UserPostgres) GetUserByID(id int) (*model.ResponseUser, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("GetUserByID: can not starts transaction:%s", err)
		return nil, fmt.Errorf("getUserByID: can not starts transaction:%w", err)
	}
	var user model.ResponseUser
	result := transaction.QueryRow("SELECT id, email, created_at FROM users WHERE id = $1", id)
	if err := result.Scan(&user.ID, &user.Email, &user.CreatedAt); err != nil {
		u.logger.Errorf("GetUserByID: error while scanning for user:%s", err)
		return nil, fmt.Errorf("getUserByID: repository error:%w", err)
	}

	return &user, transaction.Commit()
}

// GetUserPasswordByID ...
func (u UserPostgres) GetUserPasswordByID(id int) (string, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("GetUserPasswordByID: can not starts transaction:%s", err)
		return "", fmt.Errorf("getUserPasswordByID: can not starts transaction:%w", err)
	}
	var password string
	result := transaction.QueryRow("SELECT password FROM users WHERE id = $1", id)
	if err := result.Scan(&password); err != nil {
		u.logger.Errorf("GetUserPasswordByID: error while scanning for user:%s", err)
		return "", fmt.Errorf("getUserPasswordByID: repository error:%w", err)
	}

	return password, transaction.Commit()
}

// GetUserAll ...
func (u *UserPostgres) GetUserAll(page int, limit int) ([]model.ResponseUser, int, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("GetUserAll: can not starts transaction:%s", err)
		return nil, 0, fmt.Errorf("getUserAll: can not starts transaction:%w", err)
	}
	var Users []model.ResponseUser
	var query string
	var pages int
	if page == 0 || limit == 0 {
		query = "SELECT id, email, created_at FROM users"
	} else {
		query = fmt.Sprintf("SELECT id, email, created_at FROM users ORDER BY id LIMIT %d OFFSET %d", limit, (page-1)*limit)
	}

	rows, err := transaction.Query(query)
	if err != nil {
		u.logger.Errorf("GetUserAll: can not executes a query:%s", err)
		return nil, 0, fmt.Errorf("getUserAll:repository error:%w", err)
	}

	for rows.Next() {
		var User model.ResponseUser
		if err := rows.Scan(&User.ID, &User.Email, &User.CreatedAt); err != nil {
			u.logger.Errorf("Error while scanning for user:%s", err)
			return nil, 0, fmt.Errorf("getUserAll:repository error:%w", err)
		}
		Users = append(Users, User)
	}
	query = "SELECT CEILING(COUNT(id)/$1::float) FROM users"
	row := transaction.QueryRow(query, limit)
	if err := row.Scan(&pages); err != nil {
		u.logger.Errorf("Error while scanning for pages:%s", err)
		return nil, 0, fmt.Errorf("getUserAll: error while scanning for pages:%w", err)
	}
	return Users, pages, transaction.Commit()
}

// CreateUser ...
func (u *UserPostgres) CreateUser(user *model.CreateUser) (int, error) {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("CreateUser: can not starts transaction:%s", err)
		return 0, fmt.Errorf("createUser: can not starts transaction:%w", err)
	}
	var id int
	defer transaction.Rollback()
	row := transaction.QueryRow("INSERT INTO users (email, password, created_at) VALUES ($1, $2, $3) RETURNING id", user.Email, user.Password, time.Now())
	if err := row.Scan(&id); err != nil {
		u.logger.Errorf("CreateUser: error while scanning for user:%s", err)
		return 0, fmt.Errorf("createUser: error while scanning for user:%w", err)
	}
	return id, transaction.Commit()
}

// UpdateUser ...
func (u *UserPostgres) UpdateUser(user *model.UpdateUser, id int) error {
	transaction, err := u.db.Begin()
	if err != nil {
		u.logger.Errorf("UpdateUser: can not starts transaction:%s", err)
		return fmt.Errorf("updateUser: can not starts transaction:%w", err)
	}
	defer transaction.Rollback()
	_, err = transaction.Exec("UPDATE users SET password =$1 WHERE id=$2", user.NewPassword, id)
	if err != nil {
		u.logger.Errorf("UpdateUser: error while updating user:%s", err)
		return fmt.Errorf("updateUser: error while updating user:%w", err)
	}
	return transaction.Commit()
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
