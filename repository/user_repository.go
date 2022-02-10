package repository

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/mail"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"strings"
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
	db := u.db

	var user model.User


	result, err := db.Query("SELECT id, email, password, created_at FROM users WHERE id = $1", id)

	if err != nil {
		// print stack trace
		log.Println("Error query user: " + err.Error())
		return nil, err
	}

	for result.Next() {

		err := result.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)

		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

// GetUserAll ...
func (u *UserPostgres) GetUserAll() ([]model.User, error) {
	db := u.db

	var User model.User
	var Users []model.User

	rows, err := db.Query("SELECT id, email, password, created_at FROM users")

	if err != nil {
		log.Println("Error query user: " + err.Error())
		return Users, err
	}

	for rows.Next() {
	if err := rows.Scan(&User.ID, &User.Email, &User.Password, &User.CreatedAt); err != nil {
			return Users, err
		}
		Users = append(Users, User)
	}

	return Users, nil
}
func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 12
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return  b.String()
}

// CreateUser ...
func (u *UserPostgres) CreateUser(user *model.User) (*model.User, error) {
	db := u.db

	userr := model.User{}

	str := GeneratePassword()
	if user.Password == ""{
		user.Password = str
	}

	userr.Password = user.Password
	userr.Email = user.Email

	hash, _ := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	user.Password = hash

	row := db.QueryRow("INSERT INTO users (email, password, created_at) VALUES ($1, $2, $3) RETURNING id, email, password, created_at", user.Email, user.Password, time.Now())
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		log.Printf("Error while scanning for user:%s", err)
		return nil, fmt.Errorf("create user: error while scanning for user:%w", err)

	}
	// send user password
	err := mail.SendEmail(userr.Email, userr.Password, "Hello, this is your personal account password")
	if err != nil{
		log.Print(err)
	}

	userr.ID = user.ID
	userr.CreatedAt = user.CreatedAt

	return &userr, nil
}

// UpdateUser ...
func (u *UserPostgres) UpdateUser(user model.User, id int) (*model.User, error) {
	db := u.db

	hash, _ := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	user.Password = hash
	row := db.QueryRow("UPDATE users SET email =$1, password =$2 WHERE id=$3 RETURNING id", user.Email, user.Password, id)
	if err := row.Scan(&user.ID); err != nil {
		log.Printf("Error while scanning for user:%s", err)
		return nil, fmt.Errorf("update user: error while scanning for user:%w", err)
	}
	// find user by id
	result, err := u.GetUserByID(user.ID)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return result, nil
}

// DeleteUserByID ...
func (u *UserPostgres) DeleteUserByID(id int) error {
	db := u.db

	res, err := u.GetUserByID(id)
	if err != nil {
		return err
	}
	if (model.User{} == *res) {
		return errors.New("no record value with id: %v" )
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
