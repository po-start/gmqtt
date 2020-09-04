// Code generated by MockGen. DO NOT EDIT.
// Source: plugin.go

// Package gmqtt is a generated GoMock package.
package gmqtt

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPlugable is a mock of Plugable interface
type MockPlugable struct {
	ctrl     *gomock.Controller
	recorder *MockPlugableMockRecorder
}

// MockPlugableMockRecorder is the mock recorder for MockPlugable
type MockPlugableMockRecorder struct {
	mock *MockPlugable
}

// NewMockPlugable creates a new mock instance
func NewMockPlugable(ctrl *gomock.Controller) *MockPlugable {
	mock := &MockPlugable{ctrl: ctrl}
	mock.recorder = &MockPlugableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlugable) EXPECT() *MockPlugableMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockPlugable) Load(service Server) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockPlugableMockRecorder) Load(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockPlugable)(nil).Load), service)
}

// Unload mocks base method
func (m *MockPlugable) Unload() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unload")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unload indicates an expected call of Unload
func (mr *MockPlugableMockRecorder) Unload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unload", reflect.TypeOf((*MockPlugable)(nil).Unload))
}

// HookWrapper mocks base method
func (m *MockPlugable) HookWrapper() HookWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HookWrapper")
	ret0, _ := ret[0].(HookWrapper)
	return ret0
}

// HookWrapper indicates an expected call of HookWrapper
func (mr *MockPlugableMockRecorder) HookWrapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HookWrapper", reflect.TypeOf((*MockPlugable)(nil).HookWrapper))
}

// Name mocks base method
func (m *MockPlugable) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockPlugableMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPlugable)(nil).Name))
}