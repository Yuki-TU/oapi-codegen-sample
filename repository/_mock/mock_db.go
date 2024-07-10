// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/db.go
//
// Generated by this command:
//
//	mockgen -source=./repository/db.go -destination=./repository/_mock/mock_db.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDBTX is a mock of DBTX interface.
type MockDBTX struct {
	ctrl     *gomock.Controller
	recorder *MockDBTXMockRecorder
}

// MockDBTXMockRecorder is the mock recorder for MockDBTX.
type MockDBTXMockRecorder struct {
	mock *MockDBTX
}

// NewMockDBTX creates a new mock instance.
func NewMockDBTX(ctrl *gomock.Controller) *MockDBTX {
	mock := &MockDBTX{ctrl: ctrl}
	mock.recorder = &MockDBTXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBTX) EXPECT() *MockDBTXMockRecorder {
	return m.recorder
}

// ExecContext mocks base method.
func (m *MockDBTX) ExecContext(arg0 context.Context, arg1 string, arg2 ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockDBTXMockRecorder) ExecContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockDBTX)(nil).ExecContext), varargs...)
}

// PrepareContext mocks base method.
func (m *MockDBTX) PrepareContext(arg0 context.Context, arg1 string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareContext", arg0, arg1)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareContext indicates an expected call of PrepareContext.
func (mr *MockDBTXMockRecorder) PrepareContext(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareContext", reflect.TypeOf((*MockDBTX)(nil).PrepareContext), arg0, arg1)
}

// QueryContext mocks base method.
func (m *MockDBTX) QueryContext(arg0 context.Context, arg1 string, arg2 ...any) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockDBTXMockRecorder) QueryContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockDBTX)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockDBTX) QueryRowContext(arg0 context.Context, arg1 string, arg2 ...any) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockDBTXMockRecorder) QueryRowContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockDBTX)(nil).QueryRowContext), varargs...)
}