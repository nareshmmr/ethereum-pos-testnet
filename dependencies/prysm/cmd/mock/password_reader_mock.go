// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/connorwstein/prysm/cmd (interfaces: PasswordReader)

// Package mock_cmd is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// PasswordReader is a mock of PasswordReader interface
type PasswordReader struct {
	ctrl     *gomock.Controller
	recorder *PasswordReaderMockRecorder
}

// PasswordReaderMockRecorder is the mock recorder for MockPasswordReader
type PasswordReaderMockRecorder struct {
	mock *PasswordReader
}

// NewPasswordReader creates a new mock instance
func NewPasswordReader(ctrl *gomock.Controller) *PasswordReader {
	mock := &PasswordReader{ctrl: ctrl}
	mock.recorder = &PasswordReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *PasswordReader) EXPECT() *PasswordReaderMockRecorder {
	return m.recorder
}

// ReadPassword mocks base method
func (m *PasswordReader) ReadPassword() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPassword")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPassword indicates an expected call of ReadPassword
func (mr *PasswordReaderMockRecorder) ReadPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPassword", reflect.TypeOf((*PasswordReader)(nil).ReadPassword))
}