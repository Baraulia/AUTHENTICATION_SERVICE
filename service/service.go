package service

import (
	auth_proto "github.com/Baraulia/AUTHENTICATION_SERVICE/GRPC"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/GRPC/grpcClient"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
)

//go:generate mockgen -source = service.go -destination = mocks/service_mock.go

type AppUser interface {
	GetUser(id int) (*model.User, error)
	GetUsers(page int, limit int) ([]model.User, error)
	CreateUser(user *model.CreateUser) (*model.User, error)
	UpdateUser(user *model.UpdateUser, id int) (int, error)
	DeleteUserByID(id int) (int, error)
	AuthUser(email string, password string) (int, error)
	GrpcExample(map[string]string) (*auth_proto.Response, error)
}

type Service struct {
	AppUser
}

func NewService(rep *repository.Repository, grpcCli *grpcClient.GRPCClient, logger logging.Logger) *Service {
	return &Service{
		AppUser: NewUserService(*rep, grpcCli, logger),
	}
}
