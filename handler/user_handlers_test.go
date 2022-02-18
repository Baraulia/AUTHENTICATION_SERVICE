package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/service"
	mock_service "stlab.itechart-group.com/go/food_delivery/authentication_service/service/mocks"
	"testing"
	"time"
)

func TestHandler_getUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, id int)

	testTable := []struct {
		name                string
		input               string
		id                  int
		mockBehavior        mockBehavior
		expectedStatusCode  int    //expected code
		expectedRequestBody string //expected response
	}{
		{
			name:  "OK",
			input: "1",
			id:    1,
			mockBehavior: func(s *mock_service.MockAppUser, id int) {
				s.EXPECT().GetUser(id).Return(&model.User{
					ID:        1,
					Email:     "test@yande.ru",
					Password:  "ghXD!36gyd",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"email":"test@yande.ru","password":"ghXD!36gyd","created_at":"2022-02-10T16:53:28.000686358Z"}`,
		},
		{
			name:                "invalid request",
			input:               "a",
			mockBehavior:        func(s *mock_service.MockAppUser, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid request"}`,
		},
		{
			name:  "non-existent id",
			input: "1",
			id:    1,
			mockBehavior: func(s *mock_service.MockAppUser, id int) {
				s.EXPECT().GetUser(id).Return(nil, fmt.Errorf("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			getUser := mock_service.NewMockAppUser(c)
			testCase.mockBehavior(getUser, testCase.id)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: getUser}
			handler := NewHandler(logger, services)

			//Init server
			r := handler.InitRoutes()

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", testCase.input), nil)

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_getUsers(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, page int, limit int)

	testTable := []struct {
		name                string
		inputQuery          string
		page                int
		limit               int
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:       "OK",
			inputQuery: "?page=1&limit=10",
			page:       1,
			limit:      10,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int) {
				s.EXPECT().GetUsers(page, limit).Return([]model.User{
					{ID: 1,
						Email:     "test@yande.ru",
						Password:  "ghXD!36gyd",
						CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
					}, {ID: 2,
						Email:     "test2@yande.ru",
						Password:  "ghgD!36gyd",
						CreatedAt: time.Date(2022, 02, 11, 16, 53, 28, 686358, time.UTC),
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"id":1,"email":"test@yande.ru","password":"ghXD!36gyd","created_at":"2022-02-10T16:53:28.000686358Z"},{"id":2,"email":"test2@yande.ru","password":"ghgD!36gyd","created_at":"2022-02-11T16:53:28.000686358Z"}]`,
		},
		{
			name:       "Empty url query",
			inputQuery: "",
			page:       0,
			limit:      0,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int) {
				s.EXPECT().GetUsers(page, limit).Return([]model.User{
					{ID: 1,
						Email:     "test@yande.ru",
						Password:  "ghXD!36gyd",
						CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
					}, {ID: 2,
						Email:     "test2@yande.ru",
						Password:  "ghgD!36gyd",
						CreatedAt: time.Date(2022, 02, 11, 16, 53, 28, 686358, time.UTC),
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"id":1,"email":"test@yande.ru","password":"ghXD!36gyd","created_at":"2022-02-10T16:53:28.000686358Z"},{"id":2,"email":"test2@yande.ru","password":"ghgD!36gyd","created_at":"2022-02-11T16:53:28.000686358Z"}]`,
		},
		{
			name:                "Invalid value of the page in url query",
			inputQuery:          "?page=a&limit=2",
			page:                0,
			limit:               0,
			mockBehavior:        func(s *mock_service.MockAppUser, page int, limit int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid url query"}`,
		},
		{
			name:                "Invalid value of the limit in url query",
			inputQuery:          "?page=1&limit=-2",
			page:                0,
			limit:               0,
			mockBehavior:        func(s *mock_service.MockAppUser, page int, limit int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid url query"}`,
		},
		{
			name:       "Server error",
			inputQuery: "?page=1&limit=10",
			page:       1,
			limit:      10,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int) {
				s.EXPECT().GetUsers(page, limit).Return(nil, fmt.Errorf("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			getUsers := mock_service.NewMockAppUser(c)
			testCase.mockBehavior(getUsers, testCase.page, testCase.limit)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: getUsers}
			handler := NewHandler(logger, services)

			//Init server
			r := gin.New()
			r.GET("/users/", handler.getUsers)

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", testCase.inputQuery), nil)

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
func TestHandler_createUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, user model.CreateUser)
	testTable := []struct {
		name                string           //the name of the test
		inputBody           string           //the body of the request
		inputUser           model.CreateUser //the structure which we send to the service
		mockBehavior        mockBehavior
		expectedStatusCode  int    //expected code
		expectedRequestBody string //expected response
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu!98Tg"}`,
			inputUser: model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
			},
			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
				s.EXPECT().CreateUser(&user).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$RxK9zHqt84USYDuQ4wYjaO6f.03rVZH5HvoDbsyxyda35cnr/3FeK",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1,"email":"test@yandex.ru","password":"$2a$10$RxK9zHqt84USYDuQ4wYjaO6f.03rVZH5HvoDbsyxyda35cnr/3FeK","created_at":"2022-02-10T16:53:28.000686358Z"}`,
		},
		{
			name:      "OK(empty password)",
			inputBody: `{"email":"test@yandex.ru"}`,
			inputUser: model.CreateUser{
				Email: "test@yandex.ru",
			},
			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
				s.EXPECT().CreateUser(&user).Return(&model.User{
					ID:        1,
					Email:     "test@yandex.ru",
					Password:  "$2a$10$RxK9zHqt84USYDuQ4wYjaO6f.03rVZH5HvoDbsyxyda35cnr/3FeK",
					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
				}, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1,"email":"test@yandex.ru","password":"$2a$10$RxK9zHqt84USYDuQ4wYjaO6f.03rVZH5HvoDbsyxyda35cnr/3FeK","created_at":"2022-02-10T16:53:28.000686358Z"}`,
		},
		{
			name:      "Invalid email",
			inputBody: `{"email":"testyandex.ru"}`,
			inputUser: model.CreateUser{
				Email: "test@yandex.ru",
			},
			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Email":"emailValidator: it is not a valid email address"}`,
		},
		{
			name:      "Invalid password",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu98Tg"}`,
			inputUser: model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "HGYKnu98Tg",
			},
			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Password":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
		},
		{
			name:      "Server error",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKn!u98Tg"}`,
			inputUser: model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "HGYKn!u98Tg",
			},
			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
				s.EXPECT().CreateUser(&user).Return(nil, errors.New("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
		{
			name:                "Empty email field",
			inputBody:           `{"password":"HGYKn!u98Tg"}`,
			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid request"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.inputUser)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: auth}
			handler := NewHandler(logger, services)

			//Init server
			r := gin.New()
			r.POST("/users/", handler.createUser)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/users/", bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
func TestHandler_updateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, user model.UpdateUser, id int)
	testTable := []struct {
		name                string //the name of the test
		inputQuery          string
		inputBody           string           //the body of the request
		inputUser           model.UpdateUser //the structure which we send to the service
		id                  int
		mockBehavior        mockBehavior
		expectedStatusCode  int    //expected code
		expectedRequestBody string //expected response
	}{
		{
			name:       "OK",
			inputQuery: "?id=1",
			inputBody:  `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
			inputUser: model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "HGYKnu!98Tg",
				NewPassword: "HGYKnu!!98Tg",
			},
			id: 1,
			mockBehavior: func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {
				s.EXPECT().UpdateUser(&user, id).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Invalid url query",
			inputQuery:          "?id=a",
			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid url query"}`,
		},
		{
			name:                "Empty one field",
			inputQuery:          "?id=1",
			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg"}`,
			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid request"}`,
		},
		{
			name:                "Invalid new password",
			inputQuery:          "?id=1",
			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu98Tg"}`,
			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"NewPassword":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
		},
		{
			name:       "Server Failure",
			inputQuery: "?id=1",
			inputBody:  `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
			inputUser: model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "HGYKnu!98Tg",
				NewPassword: "HGYKnu!!98Tg",
			},
			id: 1,
			mockBehavior: func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {
				s.EXPECT().UpdateUser(&user, id).Return(0, errors.New("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.inputUser, testCase.id)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: auth}
			handler := NewHandler(logger, services)

			//Init server
			r := gin.New()
			r.PUT("/users/", handler.updateUser)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", testCase.inputQuery), bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
func TestHandler_deleteUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, id int)
	testTable := []struct {
		name                string //the name of the test
		inputQuery          string
		inputId             string
		id                  int
		mockBehavior        mockBehavior
		expectedStatusCode  int    //expected code
		expectedRequestBody string //expected response
	}{
		{
			name:    "OK",
			inputId: "1",
			id:      1,
			mockBehavior: func(s *mock_service.MockAppUser, id int) {
				s.EXPECT().DeleteUserByID(id).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Invalid parameter",
			inputId:             "a",
			mockBehavior:        func(s *mock_service.MockAppUser, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid id"}`,
		},
		{
			name:    "Server Failure",
			inputId: "1",
			id:      1,
			mockBehavior: func(s *mock_service.MockAppUser, id int) {
				s.EXPECT().DeleteUserByID(id).Return(0, errors.New("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAppUser(c)
			testCase.mockBehavior(auth, testCase.id)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: auth}
			handler := NewHandler(logger, services)

			//Init server
			r := gin.New()
			r.DELETE("/users/:id", handler.deleteUserByID)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%s", testCase.inputId), nil)

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
