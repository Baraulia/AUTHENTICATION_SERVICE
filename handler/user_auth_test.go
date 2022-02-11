package handler

import (
	"bytes"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/service"
	mock_service "github.com/Baraulia/AUTHENTICATION_SERVICE/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_authUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAppUser, user model.AuthUser)

	testTable := []struct {
		name                string         //the name of the test
		inputBody           string         //the body of the request
		inputUser           model.AuthUser //the structure which we send to the service
		mockBehavior        mockBehavior
		expectedStatusCode  int    //expected code
		expectedRequestBody string //expected response
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu!98Tg"}`,
			inputUser: model.AuthUser{
				Email:    "test@yandex.ru",
				Password: "HGYKnu!98Tg",
			},
			mockBehavior: func(s *mock_service.MockAppUser, user model.AuthUser) {
				s.EXPECT().AuthUser(user.Email, user.Password).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
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
			r.POST("/users/login", handler.authUser)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/users/login", bytes.NewBufferString(testCase.inputBody))

			//Execute the request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
