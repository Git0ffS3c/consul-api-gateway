// Code generated by MockGen. DO NOT EDIT.
// Source: ./secrets.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	types "github.com/envoyproxy/go-control-plane/pkg/cache/types"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretCache is a mock of SecretCache interface.
type MockSecretCache struct {
	ctrl     *gomock.Controller
	recorder *MockSecretCacheMockRecorder
}

// MockSecretCacheMockRecorder is the mock recorder for MockSecretCache.
type MockSecretCacheMockRecorder struct {
	mock *MockSecretCache
}

// NewMockSecretCache creates a new mock instance.
func NewMockSecretCache(ctrl *gomock.Controller) *MockSecretCache {
	mock := &MockSecretCache{ctrl: ctrl}
	mock.recorder = &MockSecretCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretCache) EXPECT() *MockSecretCacheMockRecorder {
	return m.recorder
}

// DeleteResource mocks base method.
func (m *MockSecretCache) DeleteResource(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResource", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResource indicates an expected call of DeleteResource.
func (mr *MockSecretCacheMockRecorder) DeleteResource(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResource", reflect.TypeOf((*MockSecretCache)(nil).DeleteResource), name)
}

// UpdateResource mocks base method.
func (m *MockSecretCache) UpdateResource(name string, res types.Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResource", name, res)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResource indicates an expected call of UpdateResource.
func (mr *MockSecretCacheMockRecorder) UpdateResource(name, res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResource", reflect.TypeOf((*MockSecretCache)(nil).UpdateResource), name, res)
}

// MockSecretClient is a mock of SecretClient interface.
type MockSecretClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretClientMockRecorder
}

// MockSecretClientMockRecorder is the mock recorder for MockSecretClient.
type MockSecretClientMockRecorder struct {
	mock *MockSecretClient
}

// NewMockSecretClient creates a new mock instance.
func NewMockSecretClient(ctrl *gomock.Controller) *MockSecretClient {
	mock := &MockSecretClient{ctrl: ctrl}
	mock.recorder = &MockSecretClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretClient) EXPECT() *MockSecretClientMockRecorder {
	return m.recorder
}

// FetchSecret mocks base method.
func (m *MockSecretClient) FetchSecret(ctx context.Context, name string) (*envoy_extensions_transport_sockets_tls_v3.Secret, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSecret", ctx, name)
	ret0, _ := ret[0].(*envoy_extensions_transport_sockets_tls_v3.Secret)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FetchSecret indicates an expected call of FetchSecret.
func (mr *MockSecretClientMockRecorder) FetchSecret(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSecret", reflect.TypeOf((*MockSecretClient)(nil).FetchSecret), ctx, name)
}

// MockSecretManager is a mock of SecretManager interface.
type MockSecretManager struct {
	ctrl     *gomock.Controller
	recorder *MockSecretManagerMockRecorder
}

// MockSecretManagerMockRecorder is the mock recorder for MockSecretManager.
type MockSecretManagerMockRecorder struct {
	mock *MockSecretManager
}

// NewMockSecretManager creates a new mock instance.
func NewMockSecretManager(ctrl *gomock.Controller) *MockSecretManager {
	mock := &MockSecretManager{ctrl: ctrl}
	mock.recorder = &MockSecretManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretManager) EXPECT() *MockSecretManagerMockRecorder {
	return m.recorder
}

// Manage mocks base method.
func (m *MockSecretManager) Manage(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Manage", ctx)
}

// Manage indicates an expected call of Manage.
func (mr *MockSecretManagerMockRecorder) Manage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manage", reflect.TypeOf((*MockSecretManager)(nil).Manage), ctx)
}

// Unwatch mocks base method.
func (m *MockSecretManager) Unwatch(ctx context.Context, names []string, node string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unwatch", ctx, names, node)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unwatch indicates an expected call of Unwatch.
func (mr *MockSecretManagerMockRecorder) Unwatch(ctx, names, node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unwatch", reflect.TypeOf((*MockSecretManager)(nil).Unwatch), ctx, names, node)
}

// UnwatchAll mocks base method.
func (m *MockSecretManager) UnwatchAll(ctx context.Context, node string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnwatchAll", ctx, node)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnwatchAll indicates an expected call of UnwatchAll.
func (mr *MockSecretManagerMockRecorder) UnwatchAll(ctx, node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnwatchAll", reflect.TypeOf((*MockSecretManager)(nil).UnwatchAll), ctx, node)
}

// Watch mocks base method.
func (m *MockSecretManager) Watch(ctx context.Context, names []string, node string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", ctx, names, node)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch.
func (mr *MockSecretManagerMockRecorder) Watch(ctx, names, node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockSecretManager)(nil).Watch), ctx, names, node)
}