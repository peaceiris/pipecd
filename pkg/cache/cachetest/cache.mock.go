// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pipe-cd/pipecd/pkg/cache (interfaces: Getter,Putter,Deleter,Cache)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod --package=cachetest --destination=pkg/cache/cachetest/cache.mock.go github.com/pipe-cd/pipecd/pkg/cache Getter,Putter,Deleter,Cache
//

// Package cachetest is a generated GoMock package.
package cachetest

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGetter is a mock of Getter interface.
type MockGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGetterMockRecorder
	isgomock struct{}
}

// MockGetterMockRecorder is the mock recorder for MockGetter.
type MockGetterMockRecorder struct {
	mock *MockGetter
}

// NewMockGetter creates a new mock instance.
func NewMockGetter(ctrl *gomock.Controller) *MockGetter {
	mock := &MockGetter{ctrl: ctrl}
	mock.recorder = &MockGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetter) EXPECT() *MockGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetter) Get(key string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetterMockRecorder) Get(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetter)(nil).Get), key)
}

// GetAll mocks base method.
func (m *MockGetter) GetAll() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockGetterMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGetter)(nil).GetAll))
}

// MockPutter is a mock of Putter interface.
type MockPutter struct {
	ctrl     *gomock.Controller
	recorder *MockPutterMockRecorder
	isgomock struct{}
}

// MockPutterMockRecorder is the mock recorder for MockPutter.
type MockPutterMockRecorder struct {
	mock *MockPutter
}

// NewMockPutter creates a new mock instance.
func NewMockPutter(ctrl *gomock.Controller) *MockPutter {
	mock := &MockPutter{ctrl: ctrl}
	mock.recorder = &MockPutterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPutter) EXPECT() *MockPutterMockRecorder {
	return m.recorder
}

// Put mocks base method.
func (m *MockPutter) Put(key string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockPutterMockRecorder) Put(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockPutter)(nil).Put), key, value)
}

// MockDeleter is a mock of Deleter interface.
type MockDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDeleterMockRecorder
	isgomock struct{}
}

// MockDeleterMockRecorder is the mock recorder for MockDeleter.
type MockDeleterMockRecorder struct {
	mock *MockDeleter
}

// NewMockDeleter creates a new mock instance.
func NewMockDeleter(ctrl *gomock.Controller) *MockDeleter {
	mock := &MockDeleter{ctrl: ctrl}
	mock.recorder = &MockDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleter) EXPECT() *MockDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeleter) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeleterMockRecorder) Delete(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeleter)(nil).Delete), key)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
	isgomock struct{}
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCache) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockCache) Get(key string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), key)
}

// GetAll mocks base method.
func (m *MockCache) GetAll() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCacheMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCache)(nil).GetAll))
}

// Put mocks base method.
func (m *MockCache) Put(key string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockCacheMockRecorder) Put(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockCache)(nil).Put), key, value)
}
