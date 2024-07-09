// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=./repository/repository.go -destination=./repository/_mock/mock_repository.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPreparer is a mock of Preparer interface.
type MockPreparer struct {
	ctrl     *gomock.Controller
	recorder *MockPreparerMockRecorder
}

// MockPreparerMockRecorder is the mock recorder for MockPreparer.
type MockPreparerMockRecorder struct {
	mock *MockPreparer
}

// NewMockPreparer creates a new mock instance.
func NewMockPreparer(ctrl *gomock.Controller) *MockPreparer {
	mock := &MockPreparer{ctrl: ctrl}
	mock.recorder = &MockPreparerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPreparer) EXPECT() *MockPreparerMockRecorder {
	return m.recorder
}

// PrepareContext mocks base method.
func (m *MockPreparer) PrepareContext(arg0 context.Context, arg1 string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareContext", arg0, arg1)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareContext indicates an expected call of PrepareContext.
func (mr *MockPreparerMockRecorder) PrepareContext(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareContext", reflect.TypeOf((*MockPreparer)(nil).PrepareContext), arg0, arg1)
}

// MockExecer is a mock of Execer interface.
type MockExecer struct {
	ctrl     *gomock.Controller
	recorder *MockExecerMockRecorder
}

// MockExecerMockRecorder is the mock recorder for MockExecer.
type MockExecerMockRecorder struct {
	mock *MockExecer
}

// NewMockExecer creates a new mock instance.
func NewMockExecer(ctrl *gomock.Controller) *MockExecer {
	mock := &MockExecer{ctrl: ctrl}
	mock.recorder = &MockExecerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecer) EXPECT() *MockExecerMockRecorder {
	return m.recorder
}

// ExecContext mocks base method.
func (m *MockExecer) ExecContext(arg0 context.Context, arg1 string, arg2 ...any) (sql.Result, error) {
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
func (mr *MockExecerMockRecorder) ExecContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockExecer)(nil).ExecContext), varargs...)
}

// MockQueryer is a mock of Queryer interface.
type MockQueryer struct {
	ctrl     *gomock.Controller
	recorder *MockQueryerMockRecorder
}

// MockQueryerMockRecorder is the mock recorder for MockQueryer.
type MockQueryerMockRecorder struct {
	mock *MockQueryer
}

// NewMockQueryer creates a new mock instance.
func NewMockQueryer(ctrl *gomock.Controller) *MockQueryer {
	mock := &MockQueryer{ctrl: ctrl}
	mock.recorder = &MockQueryerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryer) EXPECT() *MockQueryerMockRecorder {
	return m.recorder
}

// QueryContext mocks base method.
func (m *MockQueryer) QueryContext(arg0 context.Context, arg1 string, arg2 ...any) (*sql.Rows, error) {
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
func (mr *MockQueryerMockRecorder) QueryContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockQueryer)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockQueryer) QueryRowContext(arg0 context.Context, arg1 string, arg2 ...any) *sql.Row {
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
func (mr *MockQueryerMockRecorder) QueryRowContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockQueryer)(nil).QueryRowContext), varargs...)
}

// MockBeginner is a mock of Beginner interface.
type MockBeginner struct {
	ctrl     *gomock.Controller
	recorder *MockBeginnerMockRecorder
}

// MockBeginnerMockRecorder is the mock recorder for MockBeginner.
type MockBeginnerMockRecorder struct {
	mock *MockBeginner
}

// NewMockBeginner creates a new mock instance.
func NewMockBeginner(ctrl *gomock.Controller) *MockBeginner {
	mock := &MockBeginner{ctrl: ctrl}
	mock.recorder = &MockBeginnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeginner) EXPECT() *MockBeginnerMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockBeginner) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, opts)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockBeginnerMockRecorder) BeginTx(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockBeginner)(nil).BeginTx), ctx, opts)
}

// MockDBer is a mock of DBer interface.
type MockDBer struct {
	ctrl     *gomock.Controller
	recorder *MockDBerMockRecorder
}

// MockDBerMockRecorder is the mock recorder for MockDBer.
type MockDBerMockRecorder struct {
	mock *MockDBer
}

// NewMockDBer creates a new mock instance.
func NewMockDBer(ctrl *gomock.Controller) *MockDBer {
	mock := &MockDBer{ctrl: ctrl}
	mock.recorder = &MockDBerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBer) EXPECT() *MockDBerMockRecorder {
	return m.recorder
}

// ExecContext mocks base method.
func (m *MockDBer) ExecContext(arg0 context.Context, arg1 string, arg2 ...any) (sql.Result, error) {
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
func (mr *MockDBerMockRecorder) ExecContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockDBer)(nil).ExecContext), varargs...)
}

// PrepareContext mocks base method.
func (m *MockDBer) PrepareContext(arg0 context.Context, arg1 string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareContext", arg0, arg1)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareContext indicates an expected call of PrepareContext.
func (mr *MockDBerMockRecorder) PrepareContext(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareContext", reflect.TypeOf((*MockDBer)(nil).PrepareContext), arg0, arg1)
}

// QueryContext mocks base method.
func (m *MockDBer) QueryContext(arg0 context.Context, arg1 string, arg2 ...any) (*sql.Rows, error) {
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
func (mr *MockDBerMockRecorder) QueryContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockDBer)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockDBer) QueryRowContext(arg0 context.Context, arg1 string, arg2 ...any) *sql.Row {
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
func (mr *MockDBerMockRecorder) QueryRowContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockDBer)(nil).QueryRowContext), varargs...)
}
