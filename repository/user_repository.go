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

	result, err := db.Query("SELECT id, email, password, activated, created_at FROM users WHERE id = $1", id)
	if err != nil {
		// print stack trace
		log.Println("Error query user: " + err.Error())
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&user.ID, &user.Email, &user.Password, &user.Activated, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(user)

	return &user, nil
}

// GetUserAll ...
func (u *UserPostgres) GetUserAll() ([]model.User, error) {
	db := u.db

	var User model.User
	var Users []model.User

	rows, err := db.Query("SELECT id, email, password, activated, created_at FROM users")
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return Users, err
	}

	for rows.Next() {
		if err := rows.Scan(&User.ID, &User.Email, &User.Password, &User.Activated, &User.CreatedAt); err != nil {
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

	str := GeneratePassword()
	if user.Password == ""{
		user.Password = str
	}

	hash, _ := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	user.Password = hash

	row := db.QueryRow("INSERT INTO users (email, password, activated, created_at ) VALUES ($1, $2, $3, $4) RETURNING id", user.Email, user.Password, user.Activated, time.Now())
	if err := row.Scan(&user.ID); err != nil {
		log.Printf("Error while scanning for author:%s", err)
		return nil, fmt.Errorf("createAuthor: error while scanning for author:%w", err)
	}
	// send user password
	email := mail.NewEmail(user.Email, user.Password, "Hello, this is your personal account password")
	err := mail.SendEmail(email)
	if err != nil{
		log.Print(err)
	}

	// find user by id
	result, err := u.GetUserByID(user.ID)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return result, nil
}

// UpdateUser ...
func (u *UserPostgres) UpdateUser(user model.User, id int) (*model.User, error) {
	db := u.db

	hash, _ := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	user.Password = hash

	row := db.QueryRow("UPDATE users SET email =$1, password =$2, activated =$3 WHERE id=$4 RETURNING id", user.Email, user.Password, user.Activated, id)
	if err := row.Scan(&user.ID); err != nil {
		log.Printf("Error while scanning for author:%s", err)
		return nil, fmt.Errorf("createAuthor: error while scanning for author:%w", err)
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
	result, err := db.Query("SELECT id, email, password, activated, created_at FROM users WHERE email = $1", email)
	if err != nil {
		return User, err
	}

	for result.Next() {
		err := result.Scan(&User.ID, &User.Email, &User.Password, &User.Activated, &User.CreatedAt)
		if err != nil {
			return User, err
		}
	}

	return User, nil
}
