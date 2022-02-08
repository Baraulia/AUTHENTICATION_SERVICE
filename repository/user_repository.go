package repository

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
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
func (u *UserPostgres) GetUserByID(id int) (*model.User, error) {
	db := u.db

	var user *model.User

	result, err := db.Query("SELECT id, email, password, activated, created_at, updated_at FROM users WHERE id = $1", id)
	if err != nil {
		// print stack trace
		log.Println("Error query user: " + err.Error())
		return user, err
	}

	for result.Next() {
		err := result.Scan(&user.ID, &user.Email, &user.Password, &user.Activated, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

// GetUserAll ...
func (u *UserPostgres) GetUserAll() ([]model.User, error) {
	db := u.db

	var User model.User
	var Users []model.User

	rows, err := db.Query("SELECT id, email, password, activated, created_at, updated_at FROM users")
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return Users, err
	}

	for rows.Next() {
		if err := rows.Scan(&User.ID, &User.Email, &User.Password, &User.Activated,
			&User.CreatedAt, &User.UpdatedAt); err != nil {
			return Users, err
		}
		Users = append(Users, User)
	}

	return Users, nil
}

// CreateUser ...
func (u *UserPostgres) CreateUser(user *model.User) (*model.User, error) {
	db := u.db

	hash, _ := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	user.Password = hash

	crt, err := db.Prepare("INSERT INTO users (email, password, activated, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, email, password, activated, created_at, updated_at")
	if err != nil {
		u.logger.Errorf("%s", err)
		return nil, fmt.Errorf("%w", err)
	}

	res, err := crt.Exec(user.Email, user.Password, user.Activated, time.Now(), time.Now())
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	user.ID = int(userID)

	// find user by id
	resval, err := u.GetUserByID(user.ID)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return resval, nil
}

// UpdateUser ...
func (u *UserPostgres) UpdateUser(id int) (*model.User, error) {
	//db := u.db
	//
	//hash, _ := utils.HashPassword(User.Password, bcrypt.DefaultCost)
	//User.Password = hash
	//
	//crt, err := db.Prepare("UPDATE users SET email =$1, password =$2, activated =$3, updated_at =$4 where id=$5")
	//if err != nil {
	//	return User, err
	//}
	//timestamp := time.Now()
	//_, queryError := crt.Exec(User.ID, User.Email, User.Password, User.Activated, timestamp)
	//if queryError != nil {
	//	return User, err
	//}
	//
	//// find user by id
	//res, err := GetUserByID(User.ID)
	//if err != nil {
	//	return User, err
	//}
	db := u.db

	var user *model.User

	result, err := db.Query("SELECT id, email, password, activated, created_at, updated_at FROM users WHERE id = $1", id)
	if err != nil {
		// print stack trace
		log.Println("Error query user: " + err.Error())
		return user, err
	}

	for result.Next() {
		err := result.Scan(&user.ID, &user.Email, &user.Password, &user.Activated, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

// DeleteUserByID ...
func (u *UserPostgres) DeleteUserByID(id int) error {
	db := u.db

	res, err := u.GetUserByID(id)
	if err != nil {
		return err
	}

	s := strconv.FormatInt(int64(res.ID), 10)
	if (model.User{} == *res) { //todo непонятная проверка
		return errors.New("no record value with id: %v" + s)
	}

	crt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return err
	}
	_, queryError := crt.Exec(id)
	if queryError != nil {
		return err
	}

	return nil
}

// GetUserLogin ...
func (u *UserPostgres) GetUserLogin(email string, password string) (model.User, error) {

	var User model.User
	var err error

	// find by user
	User, err = u.GetUserByEmail(email)
	if err != nil {
		return User, err
	}

	if (model.User{} == User) {
		return User, errors.New("bad credential")
	}

	var retVal = utils.CheckPasswordHash(password, User.Password)
	if retVal == false {
		return User, errors.New("wrong password")
	}

	return User, nil
}

// GetUserByEmail ...
func (u *UserPostgres) GetUserByEmail(email string) (model.User, error) {
	db := u.db

	var User model.User
	result, err := db.Query("SELECT id, email, password, activated, created_at, updated_at FROM users WHERE email = $1", email)
	if err != nil {
		return User, err
	}

	for result.Next() {
		err := result.Scan(&User.ID, &User.Email, &User.Password, &User.Activated, &User.CreatedAt, &User.UpdatedAt)
		if err != nil {
			return User, err
		}
	}

	return User, nil
}
