package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	authProto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
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
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:  "OK",
			input: "1",
			id:    1,
			mockBehavior: func(s *mock_service.MockAppUser, id int) {
				s.EXPECT().GetUser(id).Return(&model.ResponseUser{
					ID:        1,
					Email:     "test@yande.ru",
					CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
					Role:      "Courier",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"email":"test@yande.ru","created_at":"2022-03-11","role":"Courier"}`,
		},
		{
			name:                "invalid request",
			input:               "a",
			mockBehavior:        func(s *mock_service.MockAppUser, id int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid request"}`,
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
	type mockBehavior func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters)

	testTable := []struct {
		name                string
		inputQuery          string
		page                int
		limit               int
		inputFilter         *model.RequestFilters
		inputRequestBody    string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:       "OK",
			inputQuery: "?page=1&limit=10",
			page:       1,
			limit:      10,
			inputFilter: &model.RequestFilters{
				ShowDeleted: false,
				FilterData:  false,
				StartTime:   model.MyTime{},
				EndTime:     model.MyTime{},
				FilterRole:  "",
			},
			inputRequestBody: `{}`,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters) {
				s.EXPECT().GetUsers(page, limit, filter).Return([]model.ResponseUser{
					{ID: 1,
						Email:     "test@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					}, {ID: 2,
						Email:     "test2@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					},
				}, 1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"Data":[{"id":1,"email":"test@yande.ru","created_at":"2022-03-11","role":"Courier"},{"id":2,"email":"test2@yande.ru","created_at":"2022-03-11","role":"Courier"}]}`,
		},
		{
			name:       "OK with role filter",
			inputQuery: "?page=1&limit=10",
			page:       1,
			limit:      10,
			inputFilter: &model.RequestFilters{
				ShowDeleted: true,
				FilterData:  false,
				StartTime:   model.MyTime{},
				EndTime:     model.MyTime{},
				FilterRole:  "Courier",
			},
			inputRequestBody: `{"filter_role": "Courier", "show_deleted": true}`,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters) {
				s.EXPECT().GetUsers(page, limit, filter).Return([]model.ResponseUser{
					{ID: 1,
						Email:     "test@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					}, {ID: 2,
						Email:     "test2@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					},
				}, 1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"Data":[{"id":1,"email":"test@yande.ru","created_at":"2022-03-11","role":"Courier"},{"id":2,"email":"test2@yande.ru","created_at":"2022-03-11","role":"Courier"}]}`,
		},
		{
			name:       "OK with data filter",
			inputQuery: "?page=1&limit=10",
			page:       1,
			limit:      10,
			inputFilter: &model.RequestFilters{
				ShowDeleted: true,
				FilterData:  true,
				StartTime:   model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
				EndTime:     model.MyTime{},
				FilterRole:  "",
			},
			inputRequestBody: `{"filter_data": true,  "show_deleted": true,  "start_time": "2022-03-11"}`,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters) {
				s.EXPECT().GetUsers(page, limit, filter).Return([]model.ResponseUser{
					{ID: 1,
						Email:     "test@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					}, {ID: 2,
						Email:     "test2@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					},
				}, 1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"Data":[{"id":1,"email":"test@yande.ru","created_at":"2022-03-11","role":"Courier"},{"id":2,"email":"test2@yande.ru","created_at":"2022-03-11","role":"Courier"}]}`,
		},
		{
			name:       "Empty url query",
			inputQuery: "",
			page:       0,
			limit:      0,
			inputFilter: &model.RequestFilters{
				ShowDeleted: false,
				FilterData:  false,
				StartTime:   model.MyTime{},
				EndTime:     model.MyTime{},
				FilterRole:  "",
			},
			inputRequestBody: `{}`,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters) {
				s.EXPECT().GetUsers(page, limit, filter).Return([]model.ResponseUser{
					{ID: 1,
						Email:     "test@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					}, {ID: 2,
						Email:     "test2@yande.ru",
						CreatedAt: model.MyTime{Time: time.Date(2022, 03, 11, 0, 0, 0, 0, time.UTC)},
						Role:      "Courier",
					},
				}, 1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"Data":[{"id":1,"email":"test@yande.ru","created_at":"2022-03-11","role":"Courier"},{"id":2,"email":"test2@yande.ru","created_at":"2022-03-11","role":"Courier"}]}`,
		},
		{
			name:       "Invalid value of the page in url query",
			inputQuery: "?page=a&limit=2",
			page:       0,
			limit:      0,
			inputFilter: &model.RequestFilters{
				ShowDeleted: false,
				FilterData:  false,
				StartTime:   model.MyTime{},
				EndTime:     model.MyTime{},
				FilterRole:  "",
			},
			inputRequestBody:    `{}`,
			mockBehavior:        func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid url query"}`,
		},
		{
			name:       "Invalid value of the limit in url query",
			inputQuery: "?page=1&limit=-2",
			page:       0,
			limit:      0,
			inputFilter: &model.RequestFilters{
				ShowDeleted: false,
				FilterData:  false,
				StartTime:   model.MyTime{},
				EndTime:     model.MyTime{},
				FilterRole:  "",
			},
			inputRequestBody:    `{}`,
			mockBehavior:        func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid url query"}`,
		},
		{
			name:       "Server error",
			inputQuery: "?page=1&limit=10",
			page:       1,
			limit:      10,
			inputFilter: &model.RequestFilters{
				ShowDeleted: false,
				FilterData:  false,
				StartTime:   model.MyTime{},
				EndTime:     model.MyTime{},
				FilterRole:  "",
			},
			inputRequestBody: `{}`,
			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int, filter *model.RequestFilters) {
				s.EXPECT().GetUsers(page, limit, filter).Return(nil, 0, fmt.Errorf("server error"))
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
			testCase.mockBehavior(getUsers, testCase.page, testCase.limit, testCase.inputFilter)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: getUsers}
			handler := NewHandler(logger, services)

			//Init server
			r := gin.New()
			r.POST("/users/", handler.getUsers)

			//Test request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", fmt.Sprintf("/users/%s", testCase.inputQuery), bytes.NewBufferString(testCase.inputRequestBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
func TestHandler_createCustomer(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, user model.CreateCustomer)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.CreateCustomer
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test@yandex.ru", "role_id":1, "password":"HGYKnu!98Tg"}`,
			inputUser: model.CreateCustomer{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
			},
			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateCustomer) {
				s.EXPECT().CreateCustomer(&user).Return(&authProto.GeneratedTokens{
					AccessToken:  "qwerty",
					RefreshToken: "qwerty",
				}, 1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"accessToken":"qwerty","refreshToken":"qwerty"}`,
		},
		{
			name:      "OK(empty password)",
			inputBody: `{"email":"test@yandex.ru", "role_id":1}`,
			inputUser: model.CreateCustomer{
				Email: "test@yandex.ru",
			},
			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateCustomer) {
				s.EXPECT().CreateCustomer(&user).Return(&authProto.GeneratedTokens{
					AccessToken:  "qwerty",
					RefreshToken: "qwerty",
				}, 1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"accessToken":"qwerty","refreshToken":"qwerty"}`,
		},
		{
			name:      "Invalid email",
			inputBody: `{"email":"testyandex.ru", "role_id":1}`,
			inputUser: model.CreateCustomer{
				Email: "test@yandex.ru",
			},
			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateCustomer) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Email":"emailValidator: it is not a valid email address"}`,
		},
		{
			name:      "Invalid password",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu98Tg", "role_id":1}`,
			inputUser: model.CreateCustomer{
				Email:    "test@yandex.ru",
				Password: "HGYKnu98Tg",
			},
			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateCustomer) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Password":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
		},
		{
			name:      "Server error",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKn!u98Tg", "role_id":1}`,
			inputUser: model.CreateCustomer{
				Email:    "test@yandex.ru",
				Password: "HGYKn!u98Tg",
			},
			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateCustomer) {
				s.EXPECT().CreateCustomer(&user).Return(nil, 0, errors.New("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
		{
			name:                "Empty email field",
			inputBody:           `{"password":"HGYKn!u98Tg"}`,
			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateCustomer) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid request"}`,
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
			r.POST("/users/customer", handler.createCustomer)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/users/customer", bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_createStaff(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, user *model.CreateStaff)
	type mockBehaviorCheckRole func(s *mock_service.MockAppUser, role string)
	testTable := []struct {
		name                  string
		inputBody             string
		inputUser             *model.CreateStaff
		mockBehavior          mockBehavior
		mockBehaviorCheckRole mockBehaviorCheckRole
		expectedStatusCode    int
		expectedRequestBody   string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu!98Tg", "role":"Courier"}`,
			inputUser: &model.CreateStaff{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
				Role:     "Courier",
			},
			mockBehaviorCheckRole: func(s *mock_service.MockAppUser, role string) {
				s.EXPECT().CheckInputRole(role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAppUser, user *model.CreateStaff) {
				s.EXPECT().CreateStaff(user).Return(1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "OK(empty password)",
			inputBody: `{"email":"test@yandex.ru", "role":"Courier"}`,
			inputUser: &model.CreateStaff{
				Email: "test@yandex.ru",
				Role:  "Courier",
			},
			mockBehaviorCheckRole: func(s *mock_service.MockAppUser, role string) {
				s.EXPECT().CheckInputRole(role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAppUser, user *model.CreateStaff) {
				s.EXPECT().CreateStaff(user).Return(1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Invalid email",
			inputBody: `{"email":"testyandex.ru", "role":"Courier"}`,
			inputUser: &model.CreateStaff{
				Email: "test@yandex.ru",
				Role:  "Courier",
			},
			mockBehaviorCheckRole: func(s *mock_service.MockAppUser, role string) {},
			mockBehavior:          func(s *mock_service.MockAppUser, user *model.CreateStaff) {},
			expectedStatusCode:    400,
			expectedRequestBody:   `{"Email":"emailValidator: it is not a valid email address"}`,
		},
		{
			name:      "Invalid password",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu98Tg", "role":"Courier"}`,
			inputUser: &model.CreateStaff{
				Email:    "test@yandex.ru",
				Password: "HGYKnu98Tg",
				Role:     "Courier",
			},
			mockBehaviorCheckRole: func(s *mock_service.MockAppUser, role string) {},
			mockBehavior:          func(s *mock_service.MockAppUser, user *model.CreateStaff) {},
			expectedStatusCode:    400,
			expectedRequestBody:   `{"Password":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
		},
		{
			name:      "Server error",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKn!u98Tg", "role":"Courier"}`,
			inputUser: &model.CreateStaff{
				Email:    "test@yandex.ru",
				Password: "HGYKn!u98Tg",
				Role:     "Courier",
			},
			mockBehaviorCheckRole: func(s *mock_service.MockAppUser, role string) {
				s.EXPECT().CheckInputRole(role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAppUser, user *model.CreateStaff) {
				s.EXPECT().CreateStaff(user).Return(0, errors.New("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
		{
			name:      "Incorrect role in request",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKn!u98Tg", "role":"courier"}`,
			inputUser: &model.CreateStaff{
				Email:    "test@yandex.ru",
				Password: "HGYKn!u98Tg",
				Role:     "courier",
			},
			mockBehaviorCheckRole: func(s *mock_service.MockAppUser, role string) {
				s.EXPECT().CheckInputRole(role).Return(errors.New("incorrect role came from the request"))
			},
			mockBehavior:        func(s *mock_service.MockAppUser, user *model.CreateStaff) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Incorrect role came from the request"}`,
		},
		{
			name:      "Empty email field",
			inputBody: `{"password":"HGYKn!u98Tg"}`,
			inputUser: &model.CreateStaff{
				Password: "HGYKn!u98Tg",
			},
			mockBehaviorCheckRole: func(s *mock_service.MockAppUser, role string) {},
			mockBehavior:          func(s *mock_service.MockAppUser, user *model.CreateStaff) {},
			expectedStatusCode:    400,
			expectedRequestBody:   `{"message":"invalid request"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAppUser(c)
			testCase.mockBehaviorCheckRole(auth, testCase.inputUser.Role)
			testCase.mockBehavior(auth, testCase.inputUser)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: auth}
			handler := NewHandler(logger, services)

			//Init server
			r := gin.New()
			r.POST("/users/staff", handler.createStaff)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/users/staff", bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_updateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, user model.UpdateUser)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.UpdateUser
		id                  int
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
			inputUser: model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "HGYKnu!98Tg",
				NewPassword: "HGYKnu!!98Tg",
			},
			id: 1,
			mockBehavior: func(s *mock_service.MockAppUser, user model.UpdateUser) {
				s.EXPECT().UpdateUser(&user).Return(nil)
			},
			expectedStatusCode: 204,
		},
		{
			name:                "Empty one field",
			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg"}`,
			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid request"}`,
		},
		{
			name:                "Invalid new password",
			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu98Tg"}`,
			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"NewPassword":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
		},
		{
			name:      "Server Failure",
			inputBody: `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
			inputUser: model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "HGYKnu!98Tg",
				NewPassword: "HGYKnu!!98Tg",
			},
			id: 1,
			mockBehavior: func(s *mock_service.MockAppUser, user model.UpdateUser) {
				s.EXPECT().UpdateUser(&user).Return(errors.New("server error"))
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
			testCase.mockBehavior(auth, testCase.inputUser)
			logger := logging.GetLogger()
			services := &service.Service{AppUser: auth}
			handler := NewHandler(logger, services)

			//Init server
			r := gin.New()
			r.PUT("/users/", handler.updateUser)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/users/", bytes.NewBufferString(testCase.inputBody))

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
		name                string
		inputQuery          string
		inputId             string
		id                  int
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
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
