// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/pkg/listers (interfaces: NamespaceLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNamespaceLister is a mock of NamespaceLister interface.
type MockNamespaceLister struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceListerMockRecorder
}

// MockNamespaceListerMockRecorder is the mock recorder for MockNamespaceLister.
type MockNamespaceListerMockRecorder struct {
	mock *MockNamespaceLister
}

// NewMockNamespaceLister creates a new mock instance.
func NewMockNamespaceLister(ctrl *gomock.Controller) *MockNamespaceLister {
	mock := &MockNamespaceLister{ctrl: ctrl}
	mock.recorder = &MockNamespaceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceLister) EXPECT() *MockNamespaceListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockNamespaceLister) List() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockNamespaceListerMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNamespaceLister)(nil).List))
}
