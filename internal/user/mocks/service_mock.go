// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	user "noname-realtime-support-chat/internal/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(ctx context.Context, email, name, password string) (*user.DTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, email, name, password)
	ret0, _ := ret[0].(*user.DTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(ctx, email, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), ctx, email, name, password)
}

// GetFreeUser mocks base method.
func (m *MockService) GetFreeUser(ctx context.Context) (*user.DTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFreeUser", ctx)
	ret0, _ := ret[0].(*user.DTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFreeUser indicates an expected call of GetFreeUser.
func (mr *MockServiceMockRecorder) GetFreeUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFreeUser", reflect.TypeOf((*MockService)(nil).GetFreeUser), ctx)
}

// GetUserByEmail mocks base method.
func (m *MockService) GetUserByEmail(ctx context.Context, email string, withPassword bool) (*user.DTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email, withPassword)
	ret0, _ := ret[0].(*user.DTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockServiceMockRecorder) GetUserByEmail(ctx, email, withPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockService)(nil).GetUserByEmail), ctx, email, withPassword)
}

// GetUserById mocks base method.
func (m *MockService) GetUserById(ctx context.Context, id string, withPassword bool) (*user.DTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id, withPassword)
	ret0, _ := ret[0].(*user.DTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockServiceMockRecorder) GetUserById(ctx, id, withPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockService)(nil).GetUserById), ctx, id, withPassword)
}

// UpdateUser mocks base method.
func (m *MockService) UpdateUser(ctx context.Context, userDTO *user.DTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, userDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServiceMockRecorder) UpdateUser(ctx, userDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockService)(nil).UpdateUser), ctx, userDTO)
}
