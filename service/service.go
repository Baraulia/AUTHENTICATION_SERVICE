package service

import (
	auth_proto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC/grpcClient"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type AppUser interface {
	GetUser(id int) (*model.ResponseUser, error)
	GetUsers(page int, limit int) ([]model.ResponseUser, int, error)
	CreateUser(user *model.CreateUser) (*auth_proto.GeneratedTokens, error)
	UpdateUser(user *model.UpdateUser, id int) error
	DeleteUserByID(id int) (int, error)
	AuthUser(email string, password string) (*auth_proto.GeneratedTokens, error)
	GrpcExample(string) (*auth_proto.Response, error)
}

type Service struct {
	AppUser
}

func NewService(rep *repository.Repository, grpcCli *grpcClient.GRPCClient, logger logging.Logger) *Service {
	return &Service{
		AppUser: NewUserService(*rep, grpcCli, logger),
	}
}
