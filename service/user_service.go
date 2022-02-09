package service

import (
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
)

type UserService struct {
	repo   repository.Repository
	logger logging.Logger
}

func NewUserService(repo repository.Repository, logger logging.Logger) *UserService {
	return &UserService{repo: repo, logger: logger}
}

// getUserByID godoc
// @Summary show master user by id
// @Description get string by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string
// @Failure 404 {object} model.User
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/{id} [get]

func (u *UserService) GetUser(id int) (*model.User, error) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return user, nil
}

// getUsers godoc
// @Summary show list master user
// @Description get users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Failure 400 {string} string
// @Failure 404 {object} model.User
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/ [get]

func (u *UserService) GetUsers() ([]model.User, error) {
	users, err := u.repo.GetUserAll()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return users, nil
}

// createUser godoc
// @Summary create master user
// @Description add by json master user
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body model.MUser true "User ID"
// @Success 200 {object} model.MUser
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/ [post]

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	resUser, err := u.repo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return resUser, nil
}


// updateUser godoc
// @Summary update master user
// @Description update by json master user
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body model.MUser true "User ID"
// @Success 200 {object} model.MUser
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/ [put]

func (u *UserService) UpdateUser(i model.User,id int) (*model.User, error) {
	user, err := u.repo.UpdateUser(i,id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// deleteUserByID godoc
// @Summary delete a master user by id
// @Description delete user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} model.MUser
// @Failure 400 {string} string
// @Failure 404 {object} model.MUser
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/{id} [delete]

func (u *UserService) DeleteUserByID(id int) error {

	err := u.repo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
