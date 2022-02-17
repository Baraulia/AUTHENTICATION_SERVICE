package service

import (
	"context"
	"fmt"
	auth_proto "github.com/Baraulia/AUTHENTICATION_SERVICE/GRPC"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/GRPC/grpcClient"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/mail"
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
	repo    repository.Repository
	logger  logging.Logger
	grpcCli *grpcClient.GRPCClient
}

func NewUserService(repo repository.Repository, grpcCli *grpcClient.GRPCClient, logger logging.Logger) *UserService {
	return &UserService{repo: repo, grpcCli: grpcCli, logger: logger}
}

func (u *UserService) GetUser(id int) (*model.User, error) {
	user, err := u.repo.AppUser.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUsers(page int, limit int) ([]model.User, error) {
	users, err := u.repo.AppUser.GetUserAll(page, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) CreateUser(user *model.CreateUser) (*model.User, error) {
	if user.Password == "" {
		user.Password = GeneratePassword()
	}
	pas := user.Password
	hash, err := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Errorf("CreateUser: can not generate hash from password:%s", err)
		return nil, fmt.Errorf("createUser: can not generate hash from password:%w", err)
	}
	user.Password = hash
	resUser, err := u.repo.AppUser.CreateUser(user)
	if err != nil {
		return nil, err
	}
	go mail.SendEmail(u.logger, &model.Post{
		Email:    user.Email,
		Password: pas,
	})
	return resUser, nil
}

func (u *UserService) UpdateUser(user *model.UpdateUser, id int) (int, error) {
	userDb, err := u.repo.AppUser.GetUserByID(id)
	if err != nil {
		return 0, err
	}
	if utils.CheckPasswordHash(user.OldPassword, userDb.Password) {
		newHash, err := utils.HashPassword(user.NewPassword, bcrypt.DefaultCost)
		if err != nil {
			u.logger.Errorf("UpdateUser: can not generate hash from password:%s", err)
			return 0, fmt.Errorf("updateUser: can not generate hash from password:%w", err)
		}
		user.NewPassword = newHash
		userId, err := u.repo.AppUser.UpdateUser(user, id)
		if err != nil {
			return 0, err
		}
		return userId, nil
	} else {
		u.logger.Warn("wrong email or password entered")
		return 0, fmt.Errorf("wrong email or password entered")
	}
}

func (u *UserService) DeleteUserByID(id int) (int, error) {
	userId, err := u.repo.AppUser.DeleteUserByID(id)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	length := 8 + rand.Intn(7)
	var b strings.Builder
	b.WriteRune(model.PasswordUpper[rand.Intn(len(model.PasswordUpper))])
	b.WriteRune(model.PasswordNumber[rand.Intn(len(model.PasswordNumber))])
	b.WriteRune(model.PasswordLower[rand.Intn(len(model.PasswordLower))])
	b.WriteRune(model.PasswordSpecial[rand.Intn(len(model.PasswordSpecial))])
	for i := 0; i < length-4; i++ {
		b.WriteRune(model.PasswordComposition[rand.Intn(len(model.PasswordComposition))])
	}
	return b.String()
}

func (u *UserService) GrpcExample(in string) (*auth_proto.Response, error) {
	return u.grpcCli.GetUserWithRights(context.Background(), &auth_proto.Request{AccessToken: in})
}
