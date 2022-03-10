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
	var user model.ResponseUser
	result := u.db.QueryRow("SELECT id, email, role, created_at FROM users WHERE id = $1", id)
	if err := result.Scan(&user.ID, &user.Email, &user.Role, &user.CreatedAt); err != nil {
		u.logger.Errorf("GetUserByID: error while scanning for user:%s", err)
		return nil, fmt.Errorf("getUserByID: repository error:%w", err)
	}

	return &user, nil
}

// GetUserPasswordByID ...
func (u UserPostgres) GetUserPasswordByID(id int) (string, error) {
	var password string
	result := u.db.QueryRow("SELECT password FROM users WHERE id = $1", id)
	if err := result.Scan(&password); err != nil {
		u.logger.Errorf("GetUserPasswordByID: error while scanning for user:%s", err)
		return "", fmt.Errorf("getUserPasswordByID: repository error:%w", err)
	}
	return password, nil
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
	var rows *sql.Rows
	if page == 0 || limit == 0 {
		query = "SELECT id, email, role, created_at FROM users"
		rows, err = transaction.Query(query)
		if err != nil {
			u.logger.Errorf("GetUserAll: can not executes a query:%s", err)
			return nil, 0, fmt.Errorf("getUserAll:repository error:%w", err)
		}
		pages = 1
	} else {
		query = "SELECT id, email, created_at FROM users ORDER BY id LIMIT $1 OFFSET $2"
		rows, err = transaction.Query(query, limit, (page-1)*limit)
		if err != nil {
			u.logger.Errorf("GetUserAll: can not executes a query:%s", err)
			return nil, 0, fmt.Errorf("getUserAll:repository error:%w", err)
		}
	}

	for rows.Next() {
		var User model.ResponseUser
		if err := rows.Scan(&User.ID, &User.Email, &User.Role, &User.CreatedAt); err != nil {
			u.logger.Errorf("Error while scanning for user:%s", err)
			return nil, 0, fmt.Errorf("getUserAll:repository error:%w", err)
		}
		Users = append(Users, User)
	}
	if pages == 0 {
		query = "SELECT CEILING(COUNT(id)/$1::float) FROM users"
		row := transaction.QueryRow(query, limit)
		if err := row.Scan(&pages); err != nil {
			u.logger.Errorf("Error while scanning for pages:%s", err)
		}
	}
	return Users, pages, transaction.Commit()
}

// CreateStaff ...
func (u *UserPostgres) CreateStaff(user *model.CreateStaff) (int, error) {
	var id int
	row := u.db.QueryRow("INSERT INTO users (email, password, role, created_at) VALUES ($1, $2, $3, $4) RETURNING id", user.Email, user.Password, user.Role, time.Now())
	if err := row.Scan(&id); err != nil {
		u.logger.Errorf("CreateStaff: error while scanning for user:%s", err)
		return 0, fmt.Errorf("CreateStaff: error while scanning for user:%w", err)
	}
	return id, nil
}

// CreateCustomer ...
func (u *UserPostgres) CreateCustomer(user *model.CreateCustomer) (int, error) {
	var id int
	row := u.db.QueryRow("INSERT INTO users (email, password, role, created_at) VALUES ($1, $2, $3, $4) RETURNING id", user.Email, user.Password, "Authorized Customer", time.Now())
	if err := row.Scan(&id); err != nil {
		u.logger.Errorf("CreateCustomer: error while scanning for user:%s", err)
		return 0, fmt.Errorf("CreateCustomer: error while scanning for user:%w", err)
	}
	return id, nil
}

// UpdateUser ...
func (u *UserPostgres) UpdateUser(user *model.UpdateUser, id int) error {
	_, err := u.db.Exec("UPDATE users SET password =$1 WHERE id=$2", user.NewPassword, id)
	if err != nil {
		u.logger.Errorf("UpdateUser: error while updating user:%s", err)
		return fmt.Errorf("updateUser: error while updating user:%w", err)
	}
	return nil
}

// DeleteUserByID ...
func (u *UserPostgres) DeleteUserByID(id int) (int, error) {
	var userId int
	row := u.db.QueryRow("DELETE FROM users WHERE id=$1 RETURNING id", id)
	if err := row.Scan(&userId); err != nil {
		u.logger.Errorf("DeleteUserByID: error while scanning for userId:%s", err)
		return 0, fmt.Errorf("deleteUserByID: error while scanning for userId:%w", err)
	}
	return userId, nil
}

// GetUserByEmail ...
func (u *UserPostgres) GetUserByEmail(email string) (*model.User, error) {
	var User model.User
	query := "SELECT id, email, password, created_at FROM users WHERE email = $1"
	row := u.db.QueryRow(query, email)
	if err := row.Scan(&User.ID, &User.Email, &User.Password, &User.CreatedAt); err != nil {
		u.logger.Errorf("Error while scanning for user:%s", err)
		return nil, fmt.Errorf("getUserByEmail: repository error:%w", err)

	}
	return &User, nil
}
