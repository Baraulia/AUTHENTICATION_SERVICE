package service

import (
	"context"
	"fmt"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/utils"
)

func (u *UserService) AuthUser(email string, password string) (*auth_proto.GeneratedTokens, int, error) {
	userDb, err := u.repo.AppUser.GetUserByEmail(email)
	if err != nil {
		return nil, 0, err
	}
	if utils.CheckPasswordHash(password, userDb.Password) {
		tokens, err := u.grpcCli.TokenGenerationById(context.Background(), &auth_proto.User{
			UserId: int32(userDb.ID),
		})
		if err != nil {
			u.logger.Errorf("TokenGenerationById:%s", err)
			return nil, 0, fmt.Errorf("tokenGenerationById:%w", err)
		}
		return tokens, userDb.ID, nil
	} else {
		u.logger.Warn("AuthUser: wrong email or password entered")
		return nil, 0, fmt.Errorf("wrong email or password entered")
	}
}
