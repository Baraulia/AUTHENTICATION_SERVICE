package repository

import (
	_ "database/sql"
	"errors"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/database"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/utils"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

// GetUserByID ...
func GetUserByID(id int64) (model.User, error) {
	db := database.DB

	var user model.User

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
func GetUserAll() ([]model.User, error) {
	db := database.DB

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
func CreateUser(User model.User) (model.User, error) {
	db := database.DB

	var err error

	hash, _ := utils.HashPassword(User.Password, bcrypt.DefaultCost)
	User.Password = hash

	crt, err := db.Prepare("INSERT INTO users (email, password, activated, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, email, password, activated, created_at, updated_at")
	if err != nil {
		log.Panic(err)
		return User, err
	}
	timestamp := time.Now()
	res, err := crt.Exec(User.Email, User.Password, User.Activated, timestamp, timestamp)
	if err != nil {
		log.Panic(err)
		return User, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)
		return User, err
	}
	User.ID = userID

	// find user by id
	resval, err := GetUserByID(User.ID)
	if err != nil {
		log.Panic(err)
		return User, err
	}

	return resval, nil
}

// UpdateUser ...
//func UpdateUser(User model.User) (model.User, error) {
//	db := database.DB
//
//	var err error
//
//	hash, _ := utils.HashPassword(User.Password, bcrypt.DefaultCost)
//	User.Password = hash
//
//	crt, err := db.Prepare("UPDATE users SET email =$1, password =$2, activated =$3, updated_at =$4 where id=$5")
//	if err != nil {
//		return User, err
//	}
//	timestamp := time.Now()
//	_, queryError := crt.Exec(User.ID, User.Email, User.Password, User.Activated, timestamp)
//	if queryError != nil {
//		return User, err
//	}
//
//	// find user by id
//	res, err := GetUserByID(User.ID)
//	if err != nil {
//		return User, err
//	}
//
//	return res, nil
//}

// DeleteUserByID ...
func DeleteUserByID(id int64) error {
	db := database.DB

	res, err := GetUserByID(id)
	if err != nil {
		return err
	}

	s := strconv.FormatInt(res.ID, 10)
	if (model.User{} == res) {
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
func GetUserLogin(email string, password string) (model.User, error) {

	var User model.User
	var err error

	// find by user
	User, err = GetUserByEmail(email)
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
func GetUserByEmail(email string) (model.User, error) {
	db := database.DB

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