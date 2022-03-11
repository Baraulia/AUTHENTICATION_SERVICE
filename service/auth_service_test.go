package service

//func TestService_authUser(t *testing.T) {
//	type mockBehaviorGetUser func(s *mock_repository.MockAppUser, email string)
//	type mockBehaviorGetTokens func(s *mock_authProto.MockAuthClient, id int32)
//	testTable := []struct {
//		name                  string
//		inputPassword         string
//		inputEmail            string
//		mockBehaviorGetUser   mockBehaviorGetUser
//		mockBehaviorGetTokens mockBehaviorGetTokens
//		expectedId            int
//		expectedError         error
//	}{
//		{
//			name:          "OK",
//			inputPassword: "HGYKnu!98Tg",
//			inputEmail:    "test@yandex.ru",
//			mockBehaviorGetUser: func(s *mock_repository.MockAppUser, email string) {
//				s.EXPECT().GetUserByEmail(email).Return(&model.User{
//					ID:       1,
//					Email:    "test@yandex.ru",
//					Password: "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
//					Deleted:  false,
//				}, nil)
//			},
//			mockBehaviorGetTokens: func(s *mock_authProto.MockAuthClient, id int32) {
//				s.EXPECT().TokenGenerationByUserId(context.Background(), &authProto.User{
//					UserId: 1,
//					Role:   "Superadmin",
//				}).Return(&authProto.GeneratedTokens{
//					AccessToken:  "qwerty",
//					RefreshToken: "qwerty",
//				}, nil)
//			},
//			expectedId:    1,
//			expectedError: nil,
//		},
//		{
//			name:          "Wrong password",
//			inputPassword: "HGYKnu!9Tg",
//			inputEmail:    "test@yandex.ru",
//			mockBehaviorGetUser: func(s *mock_repository.MockAppUser, email string) {
//				s.EXPECT().GetUserByEmail(email).Return(&model.User{
//					ID:       1,
//					Email:    "test@yandex.ru",
//					Password: "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
//					Deleted:  false,
//				}, nil)
//			},
//			expectedError: errors.New("wrong email or password entered"),
//		},
//		{
//			name:          "Repository error",
//			inputPassword: "HGYKnu!98Tg",
//			inputEmail:    "test@yandex.ru",
//			mockBehaviorGetUser: func(s *mock_repository.MockAppUser, email string) {
//				s.EXPECT().GetUserByEmail(email).Return(nil, errors.New("repository error"))
//			},
//			expectedError: errors.New("repository error"),
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			//Init dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			auth := mock_repository.NewMockAppUser(c)
//			testCase.mockBehaviorGetUser(auth, testCase.inputEmail)
//			logger := logging.GetLogger()
//			repo := &repository.Repository{AppUser: auth}
//			grpcCli := grpcClient.NewGRPCClient("159.223.1.135")
//			service := NewService(repo, grpcCli, logger)
//			_, id, err := service.AuthUser(testCase.inputEmail, testCase.inputPassword)
//			//Assert
//			assert.Equal(t, testCase.expectedId, id)
//			assert.Equal(t, testCase.expectedError, err)
//		})
//	}
//
//}
