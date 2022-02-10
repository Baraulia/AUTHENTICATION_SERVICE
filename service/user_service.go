package service

import (
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/utils"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

type UserService struct {
	repo   repository.Repository
	logger logging.Logger
}

func NewUserService(repo repository.Repository, logger logging.Logger) *UserService {
	return &UserService{repo: repo, logger: logger}
}

func (u *UserService) GetUser(id int) (*model.User, error) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUsers(page int, limit int) ([]model.User, error) {
	users, err := u.repo.GetUserAll(page, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) CreateUser(user *model.CreateUser) (*model.User, error) {
	if user.Password == "" {
		user.Password = GeneratePassword()
	}
	resUser, err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return resUser, nil
}

func (u *UserService) UpdateUser(user *model.UpdateUser, id int) (int, error) {
	userDb, err := u.repo.AppUser.GetUserByID(id)
	if err != nil {
		return 0, err
	}
	oldHash, err := utils.HashPassword(user.OldPassword, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Errorf("UpdateUser: can not generate hash from password:%s", err)
		return 0, fmt.Errorf("updateUser: can not generate hash from password:%w", err)
	}
	if userDb.Password != oldHash {
		u.logger.Warn("wrong email or password entered")
		return 0, fmt.Errorf("wrong email or password entered")
	} else {
		newHash, err := utils.HashPassword(user.NewPassword, bcrypt.DefaultCost)
		if err != nil {
			u.logger.Errorf("UpdateUser: can not generate hash from password:%s", err)
			return 0, fmt.Errorf("updateUser: can not generate hash from password:%w", err)
		}
		user.NewPassword = newHash
		userId, err := u.repo.UpdateUser(user, id)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
}

func (u *UserService) DeleteUserByID(id int) (int, error) {
	userId, err := u.repo.DeleteUserByID(id)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(model.PasswordComposition[rand.Intn(len(model.PasswordComposition))])
	}
	return b.String()
}
