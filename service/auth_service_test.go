package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC/grpcClient"
	mock_auth_proto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC/mocks"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/repository"
	mock_repository "stlab.itechart-group.com/go/food_delivery/authentication_service/repository/mocks"
	"testing"
	"time"
)

func TestService_authUser(t *testing.T) {
	type mockBehaviorGetUser func(s *mock_repository.MockAppUser, email string)
	type mockBehaviorGetTokens func(s *mock_auth_proto.MockAuthClient, id int32)
	testTable := []struct {
		name                  string
		inputPassword         string
		inputEmail            string
		mockBehaviorGetUser   mockBehaviorGetUser
		mockBehaviorGetTokens mockBehaviorGetTokens
		expectedId            int
		expectedError         error
	}{
		{
			name:          "OK",
			inputPassword: "HGYKnu!98Tg",
			inputEmail:    "test@yandex.ru",
			mockBehaviorGetUser: func(s *mock_repository.MockAppUser, email string) {
				s.EXPECT().GetUserByEmail(email).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			mockBehaviorGetTokens: func(s *mock_auth_proto.MockAuthClient, id int32) {
				s.EXPECT().TokenGenerationById(context.Background(), &auth_proto.User{
					UserId: 1,
				}).Return(&auth_proto.GeneratedTokens{
					AccessToken:  "qwerty",
					RefreshToken: "qwerty",
				}, nil)
			},
			expectedId:    1,
			expectedError: nil,
		},
		{
			name:          "Wrong password",
			inputPassword: "HGYKnu!9Tg",
			inputEmail:    "test@yandex.ru",
			mockBehaviorGetUser: func(s *mock_repository.MockAppUser, email string) {
				s.EXPECT().GetUserByEmail(email).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedError: errors.New("wrong email or password entered"),
		},
		{
			name:          "Repository error",
			inputPassword: "HGYKnu!98Tg",
			inputEmail:    "test@yandex.ru",
			mockBehaviorGetUser: func(s *mock_repository.MockAppUser, email string) {
				s.EXPECT().GetUserByEmail(email).Return(nil, errors.New("repository error"))
			},
			expectedError: errors.New("repository error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_repository.NewMockAppUser(c)
			testCase.mockBehaviorGetUser(auth, testCase.inputEmail)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient()
			service := NewService(repo, grpcCli, logger)
			_, id, err := service.AuthUser(testCase.inputEmail, testCase.inputPassword)
			//Assert
			assert.Equal(t, testCase.expectedId, id)
			assert.Equal(t, testCase.expectedError, err)
		})
	}

}
