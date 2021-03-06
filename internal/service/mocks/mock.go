// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mebr0/tiny-url/internal/domain"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockUsers) List(ctx context.Context) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsersMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsers)(nil).List), ctx)
}

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuth) Login(ctx context.Context, toLogin domain.UserLogin) (domain.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, toLogin)
	ret0, _ := ret[0].(domain.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(ctx, toLogin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), ctx, toLogin)
}

// Register mocks base method.
func (m *MockAuth) Register(ctx context.Context, toRegister domain.UserRegister) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, toRegister)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockAuthMockRecorder) Register(ctx, toRegister interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuth)(nil).Register), ctx, toRegister)
}

// MockURLs is a mock of URLs interface.
type MockURLs struct {
	ctrl     *gomock.Controller
	recorder *MockURLsMockRecorder
}

// MockURLsMockRecorder is the mock recorder for MockURLs.
type MockURLsMockRecorder struct {
	mock *MockURLs
}

// NewMockURLs creates a new mock instance.
func NewMockURLs(ctrl *gomock.Controller) *MockURLs {
	mock := &MockURLs{ctrl: ctrl}
	mock.recorder = &MockURLsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLs) EXPECT() *MockURLsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockURLs) Create(ctx context.Context, toCreate domain.URLCreate) (domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, toCreate)
	ret0, _ := ret[0].(domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockURLsMockRecorder) Create(ctx, toCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockURLs)(nil).Create), ctx, toCreate)
}

// Delete mocks base method.
func (m *MockURLs) Delete(ctx context.Context, alias string, owner primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, alias, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockURLsMockRecorder) Delete(ctx, alias, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockURLs)(nil).Delete), ctx, alias, owner)
}

// Get mocks base method.
func (m *MockURLs) Get(ctx context.Context, alias string) (domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, alias)
	ret0, _ := ret[0].(domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockURLsMockRecorder) Get(ctx, alias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockURLs)(nil).Get), ctx, alias)
}

// GetByOwner mocks base method.
func (m *MockURLs) GetByOwner(ctx context.Context, alias string, owner primitive.ObjectID) (domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByOwner", ctx, alias, owner)
	ret0, _ := ret[0].(domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByOwner indicates an expected call of GetByOwner.
func (mr *MockURLsMockRecorder) GetByOwner(ctx, alias, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByOwner", reflect.TypeOf((*MockURLs)(nil).GetByOwner), ctx, alias, owner)
}

// ListByOwner mocks base method.
func (m *MockURLs) ListByOwner(ctx context.Context, owner primitive.ObjectID) ([]domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByOwner", ctx, owner)
	ret0, _ := ret[0].([]domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByOwner indicates an expected call of ListByOwner.
func (mr *MockURLsMockRecorder) ListByOwner(ctx, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByOwner", reflect.TypeOf((*MockURLs)(nil).ListByOwner), ctx, owner)
}

// ListByOwnerAndExpiration mocks base method.
func (m *MockURLs) ListByOwnerAndExpiration(ctx context.Context, userId primitive.ObjectID, expired bool) ([]domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByOwnerAndExpiration", ctx, userId, expired)
	ret0, _ := ret[0].([]domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByOwnerAndExpiration indicates an expected call of ListByOwnerAndExpiration.
func (mr *MockURLsMockRecorder) ListByOwnerAndExpiration(ctx, userId, expired interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByOwnerAndExpiration", reflect.TypeOf((*MockURLs)(nil).ListByOwnerAndExpiration), ctx, userId, expired)
}

// Prolong mocks base method.
func (m *MockURLs) Prolong(ctx context.Context, alias string, owner primitive.ObjectID, toProlong domain.URLProlong) (domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prolong", ctx, alias, owner, toProlong)
	ret0, _ := ret[0].(domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prolong indicates an expected call of Prolong.
func (mr *MockURLsMockRecorder) Prolong(ctx, alias, owner, toProlong interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prolong", reflect.TypeOf((*MockURLs)(nil).Prolong), ctx, alias, owner, toProlong)
}
