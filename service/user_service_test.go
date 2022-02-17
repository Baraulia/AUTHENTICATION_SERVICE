package service

import (
	"errors"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/GRPC/grpcClient"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
	mock_repository "github.com/Baraulia/AUTHENTICATION_SERVICE/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestService_GetUser(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockAppUser, id int)
	testTable := []struct {
		name          string
		inputId       int
		mockBehavior  mockBehavior
		expectedUser  *model.User
		expectedError error
	}{
		{
			name:    "OK",
			inputId: 1,
			mockBehavior: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserByID(id).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedUser: &model.User{
				ID:        1,
				Email:     "test@yandex.ru",
				Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
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
			grpcCli := grpcClient.NewGRPCClient()
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
		expectedUsers []model.User
		expectedError error
	}{
		{
			name:       "OK",
			inputPage:  1,
			inputLimit: 10,
			mockBehavior: func(s *mock_repository.MockAppUser, page int, limit int) {
				s.EXPECT().GetUserAll(page, limit).Return([]model.User{
					{ID: 1,
						Email:     "test@yande.ru",
						Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
						CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
					}, {ID: 2,
						Email:     "test2@yande.ru",
						Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
						CreatedAt: time.Date(2022, 02, 11, 16, 53, 28, 686358, time.UTC),
					},
				}, nil)
			},
			expectedUsers: []model.User{
				{ID: 1,
					Email:     "test@yande.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, {ID: 2,
					Email:     "test2@yande.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
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
				s.EXPECT().GetUserAll(page, limit).Return(nil, errors.New("repository failure"))
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
			grpcCli := grpcClient.NewGRPCClient()
			service := NewService(repo, grpcCli, logger)
			users, err := service.GetUsers(testCase.inputPage, testCase.inputLimit)
			//Assert
			assert.Equal(t, testCase.expectedUsers, users)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestService_CreateUser(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockAppUser, user *model.CreateUser)
	testTable := []struct {
		name          string
		inputUser     *model.CreateUser
		mockBehavior  mockBehavior
		expectedUser  *model.User
		expectedError error
	}{
		{
			name: "OK",
			inputUser: &model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
			},
			mockBehavior: func(s *mock_repository.MockAppUser, user *model.CreateUser) {
				s.EXPECT().CreateUser(user).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedUser: &model.User{
				ID:        1,
				Email:     "test@yandex.ru",
				Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
				CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
			},
			expectedError: nil,
		},
		{
			name: "repository error",
			inputUser: &model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
			},
			mockBehavior: func(s *mock_repository.MockAppUser, user *model.CreateUser) {
				s.EXPECT().CreateUser(user).Return(nil, errors.New("repository error"))
			},
			expectedUser:  nil,
			expectedError: errors.New("repository error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_repository.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.inputUser)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient()
			service := NewService(repo, grpcCli, logger)
			user, err := service.CreateUser(testCase.inputUser)
			//Assert
			assert.Equal(t, testCase.expectedUser, user)
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
		expectedUserId     int
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
				s.EXPECT().UpdateUser(user, id).Return(1, nil)
			},
			mockBehaviorGet: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserByID(id).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedUserId: 1,
			expectedError:  nil,
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
				s.EXPECT().GetUserByID(id).Return(nil, errors.New("error while getting user"))
			},
			expectedUserId: 0,
			expectedError:  errors.New("error while getting user"),
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
				s.EXPECT().UpdateUser(user, id).Return(0, errors.New("error while getting user"))
			},
			mockBehaviorGet: func(s *mock_repository.MockAppUser, id int) {
				s.EXPECT().GetUserByID(id).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedUserId: 0,
			expectedError:  errors.New("error while getting user"),
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
			grpcCli := grpcClient.NewGRPCClient()
			service := NewService(repo, grpcCli, logger)
			id, err := service.UpdateUser(testCase.inputUser, testCase.inputId)
			//Assert
			assert.Equal(t, testCase.expectedUserId, id)
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
			grpcCli := grpcClient.NewGRPCClient()
			service := NewService(repo, grpcCli, logger)
			id, err := service.DeleteUserByID(testCase.inputId)
			//Assert
			assert.Equal(t, testCase.expectedUserId, id)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
