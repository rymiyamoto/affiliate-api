// Code generated by MockGen. DO NOT EDIT.
// Source: rakuten.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	flag "github.com/rymiyamoto/affiliate-api/flag"
)

// MockIRakuten is a mock of IRakuten interface.
type MockIRakuten struct {
	ctrl     *gomock.Controller
	recorder *MockIRakutenMockRecorder
}

// MockIRakutenMockRecorder is the mock recorder for MockIRakuten.
type MockIRakutenMockRecorder struct {
	mock *MockIRakuten
}

// NewMockIRakuten creates a new mock instance.
func NewMockIRakuten(ctrl *gomock.Controller) *MockIRakuten {
	mock := &MockIRakuten{ctrl: ctrl}
	mock.recorder = &MockIRakutenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRakuten) EXPECT() *MockIRakutenMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockIRakuten) Exec(f *flag.Affiliate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockIRakutenMockRecorder) Exec(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockIRakuten)(nil).Exec), f)
}
