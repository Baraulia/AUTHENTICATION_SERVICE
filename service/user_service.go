package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC/grpcClient"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/mail"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/utils"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/repository"
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

func (u *UserService) GetUser(id int) (*model.ResponseUser, error) {
	user, err := u.repo.AppUser.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUsers(page int, limit int) ([]model.ResponseUser, int, error) {
	users, pages, err := u.repo.AppUser.GetUserAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return users, pages, nil
}

func (u *UserService) CreateUser(user *model.CreateUser) (*auth_proto.GeneratedTokens, int, error) {
	if user.Password == "" {
		user.Password = GeneratePassword()
	}
	pas := user.Password
	hash, err := utils.HashPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Errorf("CreateUser: can not generate hash from password:%s", err)
		return nil, 0, fmt.Errorf("createUser: can not generate hash from password:%w", err)
	}
	user.Password = hash
	id, err := u.repo.AppUser.CreateUser(user)
	if err != nil {
		return nil, 0, err
	}
	go mail.SendEmail(u.logger, &model.Post{
		Email:    user.Email,
		Password: pas,
	})
	tokens, err := u.grpcCli.TokenGenerationById(context.Background(), &auth_proto.User{
		UserId: int32(id),
		Role:   user.Role,
	})
	if err != nil {
		u.logger.Errorf("TokenGenerationById:%s", err)
		return nil, 0, fmt.Errorf("tokenGenerationById:%w", err)
	}
	return tokens, id, nil
}

func (u *UserService) UpdateUser(user *model.UpdateUser, id int) error {
	userPassword, err := u.repo.AppUser.GetUserPasswordByID(id)
	if err != nil {
		return err
	}
	if utils.CheckPasswordHash(user.OldPassword, userPassword) {
		newHash, err := utils.HashPassword(user.NewPassword, bcrypt.DefaultCost)
		if err != nil {
			u.logger.Errorf("UpdateUser: can not generate hash from password:%s", err)
			return fmt.Errorf("updateUser: can not generate hash from password:%w", err)
		}
		user.NewPassword = newHash
		err = u.repo.AppUser.UpdateUser(user, id)
		if err != nil {
			return err
		}
		return nil
	} else {
		u.logger.Warn("wrong email or password entered")
		return fmt.Errorf("wrong email or password entered")
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

func (u *UserService) GrpcExample(in string) (*auth_proto.Result, error) {

	return u.grpcCli.CheckToken(context.Background(), &auth_proto.AccessToken{AccessToken: in})
}
