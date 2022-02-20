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

func TestService_GetUser(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockAppUser, id int)
	testTable := []struct {
		name          string
		inputId       int
		mockBehavior  mockBehavior
		expectedUser  *model.ResponseUser
		expectedError error
	}{
		{
			name:    "OK",
			inputId: 1,
			mockBehavior: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserByID(id).Return(&model.ResponseUser{
					ID:        1,
					Email:     "test@yandex.ru",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedUser: &model.ResponseUser{
				ID:        1,
				Email:     "test@yandex.ru",
				CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
			},
			expectedError: nil,
		},
		{
			name:    "Repository failure",
			inputId: 1,
			mockBehavior: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserByID(id).Return(nil, errors.New("repository failure"))
			},
			expectedUser:  nil,
			expectedError: errors.New("repository failure"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_repository.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.inputId)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient("159.223.1.135")
			service := NewService(repo, grpcCli, logger)
			user, err := service.GetUser(testCase.inputId)
			//Assert
			assert.Equal(t, testCase.expectedUser, user)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestService_GetUsers(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockAppUser, page int, limit int)
	testTable := []struct {
		name          string
		inputPage     int
		inputLimit    int
		mockBehavior  mockBehavior
		expectedUsers []model.ResponseUser
		expectedError error
	}{
		{
			name:       "OK",
			inputPage:  1,
			inputLimit: 10,
			mockBehavior: func(s *mock_repository.MockAppUser, page int, limit int) {
				s.EXPECT().GetUserAll(page, limit).Return([]model.ResponseUser{
					{ID: 1,
						Email:     "test@yande.ru",
						CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
					}, {ID: 2,
						Email:     "test2@yande.ru",
						CreatedAt: time.Date(2022, 02, 11, 16, 53, 28, 686358, time.UTC),
					},
				}, 1, nil)
			},
			expectedUsers: []model.ResponseUser{
				{ID: 1,
					Email:     "test@yande.ru",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, {ID: 2,
					Email:     "test2@yande.ru",
					CreatedAt: time.Date(2022, 02, 11, 16, 53, 28, 686358, time.UTC),
				},
			},
			expectedError: nil,
		},
		{
			name:       "Repository failure",
			inputPage:  1,
			inputLimit: 10,
			mockBehavior: func(s *mock_repository.MockAppUser, page int, limit int) {
				s.EXPECT().GetUserAll(page, limit).Return(nil, 0, errors.New("repository failure"))
			},
			expectedUsers: nil,
			expectedError: errors.New("repository failure"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_repository.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.inputPage, testCase.inputLimit)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient("159.223.1.135")
			service := NewService(repo, grpcCli, logger)
			users, _, err := service.GetUsers(testCase.inputPage, testCase.inputLimit)
			//Assert
			assert.Equal(t, testCase.expectedUsers, users)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestService_CreateUser(t *testing.T) {
	type mockBehaviorId func(s *mock_repository.MockAppUser, user *model.CreateUser)
	type mockBehaviorGetTokens func(s *mock_auth_proto.MockAuthClient, id int32)
	testTable := []struct {
		name                  string
		inputUser             *model.CreateUser
		mockBehaviorId        mockBehaviorId
		mockBehaviorGetTokens mockBehaviorGetTokens
		expectedError         error
	}{
		{
			name: "OK",
			inputUser: &model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
			},
			mockBehaviorId: func(s *mock_repository.MockAppUser, user *model.CreateUser) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			mockBehaviorGetTokens: func(s *mock_auth_proto.MockAuthClient, id int32) {
				s.EXPECT().TokenGenerationById(context.Background(), &auth_proto.User{
					UserId: 1,
				}).Return(&auth_proto.GeneratedTokens{
					AccessToken:  "qwerty",
					RefreshToken: "qwerty",
				}, nil)
			},
			expectedError: nil,
		},
		{
			name: "repository error",
			inputUser: &model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
			},
			mockBehaviorId: func(s *mock_repository.MockAppUser, user *model.CreateUser) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("repository error"))
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
			testCase.mockBehaviorId(auth, testCase.inputUser)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient("159.223.1.135")
			service := NewService(repo, grpcCli, logger)
			_, _, err := service.CreateUser(testCase.inputUser)
			//Assert
			assert.Equal(t, testCase.expectedError, err)
		})
	}

}

func TestService_UpdateUser(t *testing.T) {
	type mockBehaviorUpdate func(s *mock_repository.MockAppUser, user *model.UpdateUser, id int)
	type mockBehaviorGet func(s *mock_repository.MockAppUser, id int)
	testTable := []struct {
		name               string
		inputUser          *model.UpdateUser
		inputId            int
		mockBehaviorUpdate mockBehaviorUpdate
		mockBehaviorGet    mockBehaviorGet
		expectedError      error
	}{
		{
			name: "OK",
			inputUser: &model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "HGYKnu!98Tg",
				NewPassword: "HYKnu!98Tg",
			},
			inputId: 1,
			mockBehaviorUpdate: func(s *mock_repository.MockAppUser, user *model.UpdateUser, id int) {
				s.EXPECT().UpdateUser(user, id).Return(nil)
			},
			mockBehaviorGet: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserPasswordByID(id).Return("$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy", nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while getting user",
			inputUser: &model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "HGYKnu!98Tg",
				NewPassword: "HYKnu!98Tg",
			},
			inputId:            1,
			mockBehaviorUpdate: func(s *mock_repository.MockAppUser, user *model.UpdateUser, id int) {},
			mockBehaviorGet: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserPasswordByID(id).Return("", errors.New("error while getting user"))
			},
			expectedError: errors.New("error while getting user"),
		},
		{
			name: "Error while updating user",
			inputUser: &model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "HGYKnu!98Tg",
				NewPassword: "HYKnu!98Tg",
			},
			inputId: 1,
			mockBehaviorUpdate: func(s *mock_repository.MockAppUser, user *model.UpdateUser, id int) {
				s.EXPECT().UpdateUser(user, id).Return(errors.New("error while getting user"))
			},
			mockBehaviorGet: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserPasswordByID(id).Return("$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy", nil)
			},

			expectedError: errors.New("error while getting user"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_repository.NewMockAppUser(c)
			testCase.mockBehaviorUpdate(auth, testCase.inputUser, testCase.inputId)
			testCase.mockBehaviorGet(auth, testCase.inputId)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient("159.223.1.135")
			service := NewService(repo, grpcCli, logger)
			err := service.UpdateUser(testCase.inputUser, testCase.inputId)
			//Assert

			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestService_DeleteUser(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockAppUser, id int)
	testTable := []struct {
		name           string
		inputId        int
		mockBehavior   mockBehavior
		expectedUserId int
		expectedError  error
	}{
		{
			name:    "OK",
			inputId: 1,
			mockBehavior: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().DeleteUserByID(id).Return(1, nil)
			},
			expectedUserId: 1,
			expectedError:  nil,
		},
		{
			name:    "Repository failure",
			inputId: 1,
			mockBehavior: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().DeleteUserByID(id).Return(0, errors.New("repository failure"))
			},
			expectedUserId: 0,
			expectedError:  errors.New("repository failure"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_repository.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.inputId)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient("159.223.1.135")
			service := NewService(repo, grpcCli, logger)
			id, err := service.DeleteUserByID(testCase.inputId)
			//Assert
			assert.Equal(t, testCase.expectedUserId, id)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
