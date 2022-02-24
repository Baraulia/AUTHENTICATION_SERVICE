// Code generated by MockGen. DO NOT EDIT.
// Source: auth_grpc.pb.go

// Package mock_authProto is a generated GoMock package.
package mock_authProto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	authProto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
)

// MockAuthClient is a mock of AuthClient interface.
type MockAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthClientMockRecorder
}

// MockAuthClientMockRecorder is the mock recorder for MockAuthClient.
type MockAuthClientMockRecorder struct {
	mock *MockAuthClient
}

// NewMockAuthClient creates a new mock instance.
func NewMockAuthClient(ctrl *gomock.Controller) *MockAuthClient {
	mock := &MockAuthClient{ctrl: ctrl}
	mock.recorder = &MockAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthClient) EXPECT() *MockAuthClientMockRecorder {
	return m.recorder
}

// BindUserAndRole mocks base method.
func (m *MockAuthClient) BindUserAndRole(ctx context.Context, in *authProto.User, opts ...grpc.CallOption) (*authProto.ResultBinding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BindUserAndRole", varargs...)
	ret0, _ := ret[0].(*authProto.ResultBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindUserAndRole indicates an expected call of BindUserAndRole.
func (mr *MockAuthClientMockRecorder) BindUserAndRole(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindUserAndRole", reflect.TypeOf((*MockAuthClient)(nil).BindUserAndRole), varargs...)
}

// CheckToken mocks base method.
func (m *MockAuthClient) CheckToken(ctx context.Context, in *authProto.AccessToken, opts ...grpc.CallOption) (*authProto.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckToken", varargs...)
	ret0, _ := ret[0].(*authProto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken.
func (mr *MockAuthClientMockRecorder) CheckToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockAuthClient)(nil).CheckToken), varargs...)
}

// GetSalt mocks base method.
func (m *MockAuthClient) GetSalt(ctx context.Context, in *authProto.ReqSalt, opts ...grpc.CallOption) (*authProto.Salt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSalt", varargs...)
	ret0, _ := ret[0].(*authProto.Salt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSalt indicates an expected call of GetSalt.
func (mr *MockAuthClientMockRecorder) GetSalt(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSalt", reflect.TypeOf((*MockAuthClient)(nil).GetSalt), varargs...)
}

// GetUserWithRights mocks base method.
func (m *MockAuthClient) GetUserWithRights(ctx context.Context, in *authProto.AccessToken, opts ...grpc.CallOption) (*authProto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserWithRights", varargs...)
	ret0, _ := ret[0].(*authProto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithRights indicates an expected call of GetUserWithRights.
func (mr *MockAuthClientMockRecorder) GetUserWithRights(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithRights", reflect.TypeOf((*MockAuthClient)(nil).GetUserWithRights), varargs...)
}

// TokenGenerationById mocks base method.
func (m *MockAuthClient) TokenGenerationById(ctx context.Context, in *authProto.User, opts ...grpc.CallOption) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TokenGenerationById", varargs...)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenGenerationById indicates an expected call of TokenGenerationById.
func (mr *MockAuthClientMockRecorder) TokenGenerationById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenGenerationById", reflect.TypeOf((*MockAuthClient)(nil).TokenGenerationById), varargs...)
}

// TokenGenerationByRefresh mocks base method.
func (m *MockAuthClient) TokenGenerationByRefresh(ctx context.Context, in *authProto.RefreshToken, opts ...grpc.CallOption) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TokenGenerationByRefresh", varargs...)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenGenerationByRefresh indicates an expected call of TokenGenerationByRefresh.
func (mr *MockAuthClientMockRecorder) TokenGenerationByRefresh(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenGenerationByRefresh", reflect.TypeOf((*MockAuthClient)(nil).TokenGenerationByRefresh), varargs...)
}

// MockAuthServer is a mock of AuthServer interface.
type MockAuthServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServerMockRecorder
}

// MockAuthServerMockRecorder is the mock recorder for MockAuthServer.
type MockAuthServerMockRecorder struct {
	mock *MockAuthServer
}

// NewMockAuthServer creates a new mock instance.
func NewMockAuthServer(ctrl *gomock.Controller) *MockAuthServer {
	mock := &MockAuthServer{ctrl: ctrl}
	mock.recorder = &MockAuthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServer) EXPECT() *MockAuthServerMockRecorder {
	return m.recorder
}

// BindUserAndRole mocks base method.
func (m *MockAuthServer) BindUserAndRole(arg0 context.Context, arg1 *authProto.User) (*authProto.ResultBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindUserAndRole", arg0, arg1)
	ret0, _ := ret[0].(*authProto.ResultBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindUserAndRole indicates an expected call of BindUserAndRole.
func (mr *MockAuthServerMockRecorder) BindUserAndRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindUserAndRole", reflect.TypeOf((*MockAuthServer)(nil).BindUserAndRole), arg0, arg1)
}

// CheckToken mocks base method.
func (m *MockAuthServer) CheckToken(arg0 context.Context, arg1 *authProto.AccessToken) (*authProto.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckToken", arg0, arg1)
	ret0, _ := ret[0].(*authProto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken.
func (mr *MockAuthServerMockRecorder) CheckToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockAuthServer)(nil).CheckToken), arg0, arg1)
}

// GetSalt mocks base method.
func (m *MockAuthServer) GetSalt(arg0 context.Context, arg1 *authProto.ReqSalt) (*authProto.Salt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSalt", arg0, arg1)
	ret0, _ := ret[0].(*authProto.Salt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSalt indicates an expected call of GetSalt.
func (mr *MockAuthServerMockRecorder) GetSalt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSalt", reflect.TypeOf((*MockAuthServer)(nil).GetSalt), arg0, arg1)
}

// GetUserWithRights mocks base method.
func (m *MockAuthServer) GetUserWithRights(arg0 context.Context, arg1 *authProto.AccessToken) (*authProto.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWithRights", arg0, arg1)
	ret0, _ := ret[0].(*authProto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithRights indicates an expected call of GetUserWithRights.
func (mr *MockAuthServerMockRecorder) GetUserWithRights(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithRights", reflect.TypeOf((*MockAuthServer)(nil).GetUserWithRights), arg0, arg1)
}

// TokenGenerationById mocks base method.
func (m *MockAuthServer) TokenGenerationById(arg0 context.Context, arg1 *authProto.User) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenGenerationById", arg0, arg1)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenGenerationById indicates an expected call of TokenGenerationById.
func (mr *MockAuthServerMockRecorder) TokenGenerationById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenGenerationById", reflect.TypeOf((*MockAuthServer)(nil).TokenGenerationById), arg0, arg1)
}

// TokenGenerationByRefresh mocks base method.
func (m *MockAuthServer) TokenGenerationByRefresh(arg0 context.Context, arg1 *authProto.RefreshToken) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenGenerationByRefresh", arg0, arg1)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenGenerationByRefresh indicates an expected call of TokenGenerationByRefresh.
func (mr *MockAuthServerMockRecorder) TokenGenerationByRefresh(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenGenerationByRefresh", reflect.TypeOf((*MockAuthServer)(nil).TokenGenerationByRefresh), arg0, arg1)
}

// mustEmbedUnimplementedAuthServer mocks base method.
func (m *MockAuthServer) mustEmbedUnimplementedAuthServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServer")
}

// mustEmbedUnimplementedAuthServer indicates an expected call of mustEmbedUnimplementedAuthServer.
func (mr *MockAuthServerMockRecorder) mustEmbedUnimplementedAuthServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServer", reflect.TypeOf((*MockAuthServer)(nil).mustEmbedUnimplementedAuthServer))
}

// MockUnsafeAuthServer is a mock of UnsafeAuthServer interface.
type MockUnsafeAuthServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAuthServerMockRecorder
}

// MockUnsafeAuthServerMockRecorder is the mock recorder for MockUnsafeAuthServer.
type MockUnsafeAuthServerMockRecorder struct {
	mock *MockUnsafeAuthServer
}

// NewMockUnsafeAuthServer creates a new mock instance.
func NewMockUnsafeAuthServer(ctrl *gomock.Controller) *MockUnsafeAuthServer {
	mock := &MockUnsafeAuthServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAuthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAuthServer) EXPECT() *MockUnsafeAuthServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAuthServer mocks base method.
func (m *MockUnsafeAuthServer) mustEmbedUnimplementedAuthServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServer")
}

// mustEmbedUnimplementedAuthServer indicates an expected call of mustEmbedUnimplementedAuthServer.
func (mr *MockUnsafeAuthServerMockRecorder) mustEmbedUnimplementedAuthServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServer", reflect.TypeOf((*MockUnsafeAuthServer)(nil).mustEmbedUnimplementedAuthServer))
}
