package grpcClient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
)

var logger = logging.GetLogger()

type GRPCClient struct {
	cli auth_proto.AuthClient
}

func NewGRPCClient() *GRPCClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:8090", os.Getenv("HOST")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("NewGRPCClient, Dial:%s", err)
	}
	cli := auth_proto.NewAuthClient(conn)
	return &GRPCClient{cli: cli}
}

func (c *GRPCClient) GetUserWithRights(ctx context.Context, in *auth_proto.Request, opts ...grpc.CallOption) (*auth_proto.Response, error) {
	return c.cli.GetUserWithRights(ctx, in)
}

func (c *GRPCClient) CheckToken(ctx context.Context, in *auth_proto.AccessToken, opts ...grpc.CallOption) (*auth_proto.Result, error) {
	return c.cli.CheckToken(ctx, in)
}

func (c *GRPCClient) TokenGenerationByRefresh(ctx context.Context, in *auth_proto.RefreshToken, opts ...grpc.CallOption) (*auth_proto.GeneratedTokens, error) {
	return nil, nil
}

func (c *GRPCClient) TokenGenerationById(ctx context.Context, in *auth_proto.User, opts ...grpc.CallOption) (*auth_proto.GeneratedTokens, error) {
	return c.cli.TokenGenerationById(ctx, in)
}

func (c *GRPCClient) GetSalt(ctx context.Context, in *auth_proto.ReqSalt, opts ...grpc.CallOption) (*auth_proto.Salt, error) {
	return nil, nil
}
