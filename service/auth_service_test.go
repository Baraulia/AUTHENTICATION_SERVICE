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

func TestService_authUser(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockAppUser, email string)
	testTable := []struct {
		name          string
		inputPassword string
		inputEmail    string
		mockBehavior  mockBehavior
		expectedId    int
		expectedError error
	}{
		{
			name:          "OK",
			inputPassword: "HGYKnu!98Tg",
			inputEmail:    "test@yandex.ru",
			mockBehavior: func(s *mock_repository.MockAppUser, email string) {
				s.EXPECT().GetUserByEmail(email).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedId:    1,
			expectedError: nil,
		},
		{
			name:          "Wrong password",
			inputPassword: "HGYKnu!9Tg",
			inputEmail:    "test@yandex.ru",
			mockBehavior: func(s *mock_repository.MockAppUser, email string) {
				s.EXPECT().GetUserByEmail(email).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedId:    0,
			expectedError: errors.New("wrong email or password entered"),
		},
		{
			name:          "Repository error",
			inputPassword: "HGYKnu!98Tg",
			inputEmail:    "test@yandex.ru",
			mockBehavior: func(s *mock_repository.MockAppUser, email string) {
				s.EXPECT().GetUserByEmail(email).Return(nil, errors.New("repository error"))
			},
			expectedId:    0,
			expectedError: errors.New("repository error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_repository.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.inputEmail)
			logger := logging.GetLogger()
			repo := &repository.Repository{AppUser: auth}
			grpcCli := grpcClient.NewGRPCClient()
			service := NewService(repo, grpcCli, logger)
			id, err := service.AuthUser(testCase.inputEmail, testCase.inputPassword)
			//Assert
			assert.Equal(t, testCase.expectedId, id)
			assert.Equal(t, testCase.expectedError, err)
		})
	}

}
