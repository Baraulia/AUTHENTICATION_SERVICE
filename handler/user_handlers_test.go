package handler

import (
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/service"
	mock_service "github.com/Baraulia/AUTHENTICATION_SERVICE/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
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
