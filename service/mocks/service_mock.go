// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC"
	model "stlab.itechart-group.com/go/food_delivery/authentication_service/model"
)

// MockAppUser is a mock of AppUser interface.
type MockAppUser struct {
	ctrl     *gomock.Controller
	recorder *MockAppUserMockRecorder
}

// MockAppUserMockRecorder is the mock recorder for MockAppUser.
type MockAppUserMockRecorder struct {
	mock *MockAppUser
}

// NewMockAppUser creates a new mock instance.
func NewMockAppUser(ctrl *gomock.Controller) *MockAppUser {
	mock := &MockAppUser{ctrl: ctrl}
	mock.recorder = &MockAppUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUser) EXPECT() *MockAppUserMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockAppUser) AuthUser(email, password string) (*auth_proto.GeneratedTokens, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUser", email, password)
	ret0, _ := ret[0].(*auth_proto.GeneratedTokens)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockAppUserMockRecorder) AuthUser(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockAppUser)(nil).AuthUser), email, password)
}

// CreateCustomer mocks base method.
func (m *MockAppUser) CreateCustomer(user *model.CreateUser) (*auth_proto.GeneratedTokens, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", user)
	ret0, _ := ret[0].(*auth_proto.GeneratedTokens)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockAppUserMockRecorder) CreateCustomer(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockAppUser)(nil).CreateCustomer), user)
}

// CreateStaff mocks base method.
func (m *MockAppUser) CreateStaff(user *model.CreateUser) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStaff", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStaff indicates an expected call of CreateStaff.
func (mr *MockAppUserMockRecorder) CreateStaff(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStaff", reflect.TypeOf((*MockAppUser)(nil).CreateStaff), user)
}

// DeleteUserByID mocks base method.
func (m *MockAppUser) DeleteUserByID(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserByID", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserByID indicates an expected call of DeleteUserByID.
func (mr *MockAppUserMockRecorder) DeleteUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByID", reflect.TypeOf((*MockAppUser)(nil).DeleteUserByID), id)
}

// GetUser mocks base method.
func (m *MockAppUser) GetUser(id int) (*model.ResponseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*model.ResponseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAppUserMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAppUser)(nil).GetUser), id)
}

// GetUsers mocks base method.
func (m *MockAppUser) GetUsers(page, limit int) ([]model.ResponseUser, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", page, limit)
	ret0, _ := ret[0].([]model.ResponseUser)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockAppUserMockRecorder) GetUsers(page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockAppUser)(nil).GetUsers), page, limit)
}

// GrpcExample mocks base method.
func (m *MockAppUser) GrpcExample(arg0 string) (*auth_proto.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrpcExample", arg0)
	ret0, _ := ret[0].(*auth_proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GrpcExample indicates an expected call of GrpcExample.
func (mr *MockAppUserMockRecorder) GrpcExample(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrpcExample", reflect.TypeOf((*MockAppUser)(nil).GrpcExample), arg0)
}

// UpdateUser mocks base method.
func (m *MockAppUser) UpdateUser(user *model.UpdateUser, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", user, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockAppUserMockRecorder) UpdateUser(user, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAppUser)(nil).UpdateUser), user, id)
}
