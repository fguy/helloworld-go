// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fguy/helloworld-go/repositories/helloworld (interfaces: Interface)

// Package mock_helloworld is a generated GoMock package.
package mock_helloworld

import (
	context "context"
	entities "github.com/fguy/helloworld-go/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// GetPage mocks base method
func (m *MockInterface) GetPage(arg0 context.Context, arg1 string) (*entities.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPage", arg0, arg1)
	ret0, _ := ret[0].(*entities.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage
func (mr *MockInterfaceMockRecorder) GetPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockInterface)(nil).GetPage), arg0, arg1)
}
