// Code generated by MockGen. DO NOT EDIT.
// Source: ../../internal/server/server.go
//
// Generated by this command:
//
//	mockgen -source=../../internal/server/server.go -destination ./mocks_server.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tracker "github.com/ansoncht/Cat-Food-Helper/internal/pb"
	gomock "go.uber.org/mock/gomock"
)

// MockTrackerService is a mock of TrackerService interface.
type MockTrackerService struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerServiceMockRecorder
}

// MockTrackerServiceMockRecorder is the mock recorder for MockTrackerService.
type MockTrackerServiceMockRecorder struct {
	mock *MockTrackerService
}

// NewMockTrackerService creates a new mock instance.
func NewMockTrackerService(ctrl *gomock.Controller) *MockTrackerService {
	mock := &MockTrackerService{ctrl: ctrl}
	mock.recorder = &MockTrackerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackerService) EXPECT() *MockTrackerServiceMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockTrackerService) RegisterUser(ctx context.Context, user *tracker.UserRequest) (*tracker.UserReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, user)
	ret0, _ := ret[0].(*tracker.UserReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockTrackerServiceMockRecorder) RegisterUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockTrackerService)(nil).RegisterUser), ctx, user)
}