package handler

//func TestHandler_getUser(t *testing.T) {
//	type mockBehavior func(s *mock_service.MockAppUser, id int)
//
//	testTable := []struct {
//		name                string
//		input               string
//		id                  int
//		mockBehavior        mockBehavior
//		expectedStatusCode  int
//		expectedRequestBody string
//	}{
//		{
//			name:  "OK",
//			input: "1",
//			id:    1,
//			mockBehavior: func(s *mock_service.MockAppUser, id int) {
//				s.EXPECT().GetUser(id).Return(&model.ResponseUser{
//					ID:        1,
//					Email:     "test@yande.ru",
//					CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
//				}, nil)
//			},
//			expectedStatusCode:  200,
//			expectedRequestBody: `{"id":1,"email":"test@yande.ru","created_at":"2022-02-10T16:53:28.000686358Z"}`,
//		},
//		{
//			name:                "invalid request",
//			input:               "a",
//			mockBehavior:        func(s *mock_service.MockAppUser, id int) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"invalid request"}`,
//		},
//		{
//			name:  "non-existent id",
//			input: "1",
//			id:    1,
//			mockBehavior: func(s *mock_service.MockAppUser, id int) {
//				s.EXPECT().GetUser(id).Return(nil, fmt.Errorf("server error"))
//			},
//			expectedStatusCode:  500,
//			expectedRequestBody: `{"message":"server error"}`,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			//Init dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			getUser := mock_service.NewMockAppUser(c)
//			testCase.mockBehavior(getUser, testCase.id)
//			logger := logging.GetLogger()
//			services := &service.Service{AppUser: getUser}
//			handler := NewHandler(logger, services)
//
//			//Init server
//			r := handler.InitRoutes()
//
//			//Test request
//			w := httptest.NewRecorder()
//
//			req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", testCase.input), nil)
//
//			//Execute the request
//			r.ServeHTTP(w, req)
//
//			//Assert
//			assert.Equal(t, testCase.expectedStatusCode, w.Code)
//			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
//		})
//	}
//
//}
//
//func TestHandler_getUsers(t *testing.T) {
//	type mockBehavior func(s *mock_service.MockAppUser, page int, limit int)
//
//	testTable := []struct {
//		name                string
//		inputQuery          string
//		page                int
//		limit               int
//		mockBehavior        mockBehavior
//		expectedStatusCode  int
//		expectedRequestBody string
//	}{
//		{
//			name:       "OK",
//			inputQuery: "?page=1&limit=10",
//			page:       1,
//			limit:      10,
//			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int) {
//				s.EXPECT().GetUsers(page, limit).Return([]model.ResponseUser{
//					{ID: 1,
//						Email:     "test@yande.ru",
//						CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
//					}, {ID: 2,
//						Email:     "test2@yande.ru",
//						CreatedAt: time.Date(2022, 02, 11, 16, 53, 28, 686358, time.UTC),
//					},
//				}, 1, nil)
//			},
//			expectedStatusCode:  200,
//			expectedRequestBody: `{"Data":[{"id":1,"email":"test@yande.ru","created_at":"2022-02-10T16:53:28.000686358Z"},{"id":2,"email":"test2@yande.ru","created_at":"2022-02-11T16:53:28.000686358Z"}]}`,
//		},
//		{
//			name:       "Empty url query",
//			inputQuery: "",
//			page:       0,
//			limit:      0,
//			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int) {
//				s.EXPECT().GetUsers(page, limit).Return([]model.ResponseUser{
//					{ID: 1,
//						Email:     "test@yande.ru",
//						CreatedAt: time.Date(2022, 02, 10, 16, 53, 28, 686358, time.UTC),
//					}, {ID: 2,
//						Email:     "test2@yande.ru",
//						CreatedAt: time.Date(2022, 02, 11, 16, 53, 28, 686358, time.UTC),
//					},
//				}, 1, nil)
//			},
//			expectedStatusCode:  200,
//			expectedRequestBody: `{"Data":[{"id":1,"email":"test@yande.ru","created_at":"2022-02-10T16:53:28.000686358Z"},{"id":2,"email":"test2@yande.ru","created_at":"2022-02-11T16:53:28.000686358Z"}]}`,
//		},
//		{
//			name:                "Invalid value of the page in url query",
//			inputQuery:          "?page=a&limit=2",
//			page:                0,
//			limit:               0,
//			mockBehavior:        func(s *mock_service.MockAppUser, page int, limit int) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"Invalid url query"}`,
//		},
//		{
//			name:                "Invalid value of the limit in url query",
//			inputQuery:          "?page=1&limit=-2",
//			page:                0,
//			limit:               0,
//			mockBehavior:        func(s *mock_service.MockAppUser, page int, limit int) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"Invalid url query"}`,
//		},
//		{
//			name:       "Server error",
//			inputQuery: "?page=1&limit=10",
//			page:       1,
//			limit:      10,
//			mockBehavior: func(s *mock_service.MockAppUser, page int, limit int) {
//				s.EXPECT().GetUsers(page, limit).Return(nil, 0, fmt.Errorf("server error"))
//			},
//			expectedStatusCode:  500,
//			expectedRequestBody: `{"message":"server error"}`,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			//Init dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			getUsers := mock_service.NewMockAppUser(c)
//			testCase.mockBehavior(getUsers, testCase.page, testCase.limit)
//			logger := logging.GetLogger()
//			services := &service.Service{AppUser: getUsers}
//			handler := NewHandler(logger, services)
//
//			//Init server
//			r := gin.New()
//			r.GET("/users/", handler.getUsers)
//
//			//Test request
//			w := httptest.NewRecorder()
//
//			req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", testCase.inputQuery), nil)
//
//			//Execute the request
//			r.ServeHTTP(w, req)
//
//			//Assert
//			assert.Equal(t, testCase.expectedStatusCode, w.Code)
//			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
//		})
//	}
//
//}
//func TestHandler_createCustomer(t *testing.T) {
//	type mockBehavior func(s *mock_service.MockAppUser, user model.CreateUser)
//	testTable := []struct {
//		name                string
//		inputBody           string
//		inputUser           model.CreateUser
//		mockBehavior        mockBehavior
//		expectedStatusCode  int
//		expectedRequestBody string
//	}{
//		{
//			name:      "OK",
//			inputBody: `{"email":"test@yandex.ru", "role_id":1, "password":"HGYKnu!98Tg"}`,
//			inputUser: model.CreateUser{
//				Email:    "test@yandex.ru",
//				Password: "HGYKnu!98Tg",
//				RoleId:   1,
//			},
//			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
//				s.EXPECT().CreateCustomer(&user).Return(&authProto.GeneratedTokens{
//					AccessToken:  "qwerty",
//					RefreshToken: "qwerty",
//				}, 1, nil)
//			},
//			expectedStatusCode:  201,
//			expectedRequestBody: `{"accessToken":"qwerty","refreshToken":"qwerty"}`,
//		},
//		{
//			name:      "OK(empty password)",
//			inputBody: `{"email":"test@yandex.ru", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:  "test@yandex.ru",
//				RoleId: 1,
//			},
//			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
//				s.EXPECT().CreateCustomer(&user).Return(&authProto.GeneratedTokens{
//					AccessToken:  "qwerty",
//					RefreshToken: "qwerty",
//				}, 1, nil)
//			},
//			expectedStatusCode:  201,
//			expectedRequestBody: `{"accessToken":"qwerty","refreshToken":"qwerty"}`,
//		},
//		{
//			name:      "Invalid email",
//			inputBody: `{"email":"testyandex.ru", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:  "test@yandex.ru",
//				RoleId: 1,
//			},
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"Email":"emailValidator: it is not a valid email address"}`,
//		},
//		{
//			name:      "Invalid password",
//			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu98Tg", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:    "test@yandex.ru",
//				Password: "HGYKnu98Tg",
//				RoleId:   1,
//			},
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"Password":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
//		},
//		{
//			name:      "Server error",
//			inputBody: `{"email":"test@yandex.ru", "password":"HGYKn!u98Tg", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:    "test@yandex.ru",
//				Password: "HGYKn!u98Tg",
//				RoleId:   1,
//			},
//			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
//				s.EXPECT().CreateCustomer(&user).Return(nil, 0, errors.New("server error"))
//			},
//			expectedStatusCode:  500,
//			expectedRequestBody: `{"message":"server error"}`,
//		},
//		{
//			name:                "Empty email field",
//			inputBody:           `{"password":"HGYKn!u98Tg"}`,
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"invalid request"}`,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			//Init dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			auth := mock_service.NewMockAppUser(c)
//			testCase.mockBehavior(auth, testCase.inputUser)
//			logger := logging.GetLogger()
//			services := &service.Service{AppUser: auth}
//			handler := NewHandler(logger, services)
//
//			//Init server
//			r := gin.New()
//			r.POST("/users/customer", handler.createCustomer)
//
//			//Test request
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest("POST", "/users/customer", bytes.NewBufferString(testCase.inputBody))
//
//			//Execute the request
//			r.ServeHTTP(w, req)
//
//			//Assert
//			assert.Equal(t, testCase.expectedStatusCode, w.Code)
//			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
//		})
//	}
//
//}
//func TestHandler_createStaff(t *testing.T) {
//	type mockBehavior func(s *mock_service.MockAppUser, user model.CreateUser)
//	testTable := []struct {
//		name                string
//		inputBody           string
//		inputUser           model.CreateUser
//		mockBehavior        mockBehavior
//		expectedStatusCode  int
//		expectedRequestBody string
//	}{
//		{
//			name:      "OK",
//			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu!98Tg", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:    "test@yandex.ru",
//				Password: "HGYKnu!98Tg",
//				RoleId:   1,
//			},
//			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
//				s.EXPECT().CreateStaff(&user).Return(1, nil)
//			},
//			expectedStatusCode:  201,
//			expectedRequestBody: `{"id":1}`,
//		},
//		{
//			name:      "OK(empty password)",
//			inputBody: `{"email":"test@yandex.ru", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:  "test@yandex.ru",
//				RoleId: 1,
//			},
//			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
//				s.EXPECT().CreateStaff(&user).Return(1, nil)
//			},
//			expectedStatusCode:  201,
//			expectedRequestBody: `{"id":1}`,
//		},
//		{
//			name:      "Invalid email",
//			inputBody: `{"email":"testyandex.ru", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:  "test@yandex.ru",
//				RoleId: 1,
//			},
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"Email":"emailValidator: it is not a valid email address"}`,
//		},
//		{
//			name:      "Invalid password",
//			inputBody: `{"email":"test@yandex.ru", "password":"HGYKnu98Tg", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:    "test@yandex.ru",
//				Password: "HGYKnu98Tg",
//				RoleId:   1,
//			},
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"Password":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
//		},
//		{
//			name:      "Server error",
//			inputBody: `{"email":"test@yandex.ru", "password":"HGYKn!u98Tg", "role_id":1}`,
//			inputUser: model.CreateUser{
//				Email:    "test@yandex.ru",
//				Password: "HGYKn!u98Tg",
//				RoleId:   1,
//			},
//			mockBehavior: func(s *mock_service.MockAppUser, user model.CreateUser) {
//				s.EXPECT().CreateStaff(&user).Return(0, errors.New("server error"))
//			},
//			expectedStatusCode:  500,
//			expectedRequestBody: `{"message":"server error"}`,
//		},
//		{
//			name:                "Empty email field",
//			inputBody:           `{"password":"HGYKn!u98Tg"}`,
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.CreateUser) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"invalid request"}`,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			//Init dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			auth := mock_service.NewMockAppUser(c)
//			testCase.mockBehavior(auth, testCase.inputUser)
//			logger := logging.GetLogger()
//			services := &service.Service{AppUser: auth}
//			handler := NewHandler(logger, services)
//
//			//Init server
//			r := gin.New()
//			r.POST("/users/staff", handler.createStaff)
//
//			//Test request
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest("POST", "/users/staff", bytes.NewBufferString(testCase.inputBody))
//
//			//Execute the request
//			r.ServeHTTP(w, req)
//
//			//Assert
//			assert.Equal(t, testCase.expectedStatusCode, w.Code)
//			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
//		})
//	}
//
//}
//
//func TestHandler_updateUser(t *testing.T) {
//	type mockBehavior func(s *mock_service.MockAppUser, user model.UpdateUser, id int)
//	testTable := []struct {
//		name                string
//		inputId             string
//		inputBody           string
//		inputUser           model.UpdateUser
//		id                  int
//		mockBehavior        mockBehavior
//		expectedStatusCode  int
//		expectedRequestBody string
//	}{
//		{
//			name:      "OK",
//			inputId:   "1",
//			inputBody: `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
//			inputUser: model.UpdateUser{
//				Email:       "test@yandex.ru",
//				OldPassword: "HGYKnu!98Tg",
//				NewPassword: "HGYKnu!!98Tg",
//			},
//			id: 1,
//			mockBehavior: func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {
//				s.EXPECT().UpdateUser(&user, id).Return(nil)
//			},
//			expectedStatusCode: 204,
//		},
//		{
//			name:                "Invalid parameter",
//			inputId:             "a",
//			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"Invalid id"}`,
//		},
//		{
//			name:                "Empty one field",
//			inputId:             "1",
//			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg"}`,
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"invalid request"}`,
//		},
//		{
//			name:                "Invalid new password",
//			inputId:             "1",
//			inputBody:           `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu98Tg"}`,
//			mockBehavior:        func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"NewPassword":"passwordValidator: the password must contain at least one digit(0-9), one lowercase letter(a-z), one uppercase letter(A-Z), one special character (@,#,%,\u0026,!,$)"}`,
//		},
//		{
//			name:      "Server Failure",
//			inputId:   "1",
//			inputBody: `{"email":"test@yandex.ru", "old_password":"HGYKnu!98Tg", "new_password":"HGYKnu!!98Tg"}`,
//			inputUser: model.UpdateUser{
//				Email:       "test@yandex.ru",
//				OldPassword: "HGYKnu!98Tg",
//				NewPassword: "HGYKnu!!98Tg",
//			},
//			id: 1,
//			mockBehavior: func(s *mock_service.MockAppUser, user model.UpdateUser, id int) {
//				s.EXPECT().UpdateUser(&user, id).Return(errors.New("server error"))
//			},
//			expectedStatusCode:  500,
//			expectedRequestBody: `{"message":"server error"}`,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			//Init dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			auth := mock_service.NewMockAppUser(c)
//			testCase.mockBehavior(auth, testCase.inputUser, testCase.id)
//			logger := logging.GetLogger()
//			services := &service.Service{AppUser: auth}
//			handler := NewHandler(logger, services)
//
//			//Init server
//			r := gin.New()
//			r.PUT("/users/:id", handler.updateUser)
//
//			//Test request
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", testCase.inputId), bytes.NewBufferString(testCase.inputBody))
//
//			//Execute the request
//			r.ServeHTTP(w, req)
//
//			//Assert
//			assert.Equal(t, testCase.expectedStatusCode, w.Code)
//			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
//		})
//	}
//
//}
//func TestHandler_deleteUser(t *testing.T) {
//	type mockBehavior func(s *mock_service.MockAppUser, id int)
//	testTable := []struct {
//		name                string
//		inputQuery          string
//		inputId             string
//		id                  int
//		mockBehavior        mockBehavior
//		expectedStatusCode  int
//		expectedRequestBody string
//	}{
//		{
//			name:    "OK",
//			inputId: "1",
//			id:      1,
//			mockBehavior: func(s *mock_service.MockAppUser, id int) {
//				s.EXPECT().DeleteUserByID(id).Return(1, nil)
//			},
//			expectedStatusCode:  200,
//			expectedRequestBody: `{"id":1}`,
//		},
//		{
//			name:                "Invalid parameter",
//			inputId:             "a",
//			mockBehavior:        func(s *mock_service.MockAppUser, id int) {},
//			expectedStatusCode:  400,
//			expectedRequestBody: `{"message":"Invalid id"}`,
//		},
//		{
//			name:    "Server Failure",
//			inputId: "1",
//			id:      1,
//			mockBehavior: func(s *mock_service.MockAppUser, id int) {
//				s.EXPECT().DeleteUserByID(id).Return(0, errors.New("server error"))
//			},
//			expectedStatusCode:  500,
//			expectedRequestBody: `{"message":"server error"}`,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			//Init dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			auth := mock_service.NewMockAppUser(c)
//			testCase.mockBehavior(auth, testCase.id)
//			logger := logging.GetLogger()
//			services := &service.Service{AppUser: auth}
//			handler := NewHandler(logger, services)
//
//			//Init server
//			r := gin.New()
//			r.DELETE("/users/:id", handler.deleteUserByID)
//
//			//Test request
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%s", testCase.inputId), nil)
//
//			//Execute the request
//			r.ServeHTTP(w, req)
//
//			//Assert
//			assert.Equal(t, testCase.expectedStatusCode, w.Code)
//			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
//		})
//	}
//
//}
