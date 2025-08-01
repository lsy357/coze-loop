// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/infra/fileserver (interfaces: ObjectStorage)
//
// Generated by this command:
//
//	mockgen -destination=mocks/object_storage.go -package=mocks . ObjectStorage
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	fileserver "github.com/coze-dev/coze-loop/backend/infra/fileserver"
	gomock "go.uber.org/mock/gomock"
)

// MockObjectStorage is a mock of ObjectStorage interface.
type MockObjectStorage struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStorageMockRecorder
	isgomock struct{}
}

// MockObjectStorageMockRecorder is the mock recorder for MockObjectStorage.
type MockObjectStorageMockRecorder struct {
	mock *MockObjectStorage
}

// NewMockObjectStorage creates a new mock instance.
func NewMockObjectStorage(ctrl *gomock.Controller) *MockObjectStorage {
	mock := &MockObjectStorage{ctrl: ctrl}
	mock.recorder = &MockObjectStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStorage) EXPECT() *MockObjectStorageMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockObjectStorage) Download(ctx context.Context, key string, w io.WriterAt, opts ...fileserver.DownloadOpt) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key, w}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Download", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download.
func (mr *MockObjectStorageMockRecorder) Download(ctx, key, w any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key, w}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockObjectStorage)(nil).Download), varargs...)
}

// Read mocks base method.
func (m *MockObjectStorage) Read(ctx context.Context, key string, opts ...fileserver.DownloadOpt) (fileserver.Reader, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(fileserver.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockObjectStorageMockRecorder) Read(ctx, key any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockObjectStorage)(nil).Read), varargs...)
}

// Remove mocks base method.
func (m *MockObjectStorage) Remove(ctx context.Context, key string, opts ...fileserver.RemoveOpt) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Remove", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockObjectStorageMockRecorder) Remove(ctx, key any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockObjectStorage)(nil).Remove), varargs...)
}

// SignDownloadReq mocks base method.
func (m *MockObjectStorage) SignDownloadReq(ctx context.Context, key string, opts ...fileserver.SignOpt) (string, http.Header, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignDownloadReq", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(http.Header)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignDownloadReq indicates an expected call of SignDownloadReq.
func (mr *MockObjectStorageMockRecorder) SignDownloadReq(ctx, key any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignDownloadReq", reflect.TypeOf((*MockObjectStorage)(nil).SignDownloadReq), varargs...)
}

// SignUploadReq mocks base method.
func (m *MockObjectStorage) SignUploadReq(ctx context.Context, key string, opts ...fileserver.SignOpt) (string, http.Header, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignUploadReq", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(http.Header)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignUploadReq indicates an expected call of SignUploadReq.
func (mr *MockObjectStorageMockRecorder) SignUploadReq(ctx, key any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUploadReq", reflect.TypeOf((*MockObjectStorage)(nil).SignUploadReq), varargs...)
}

// Stat mocks base method.
func (m *MockObjectStorage) Stat(ctx context.Context, key string, opts ...fileserver.StatOpt) (*fileserver.ObjectInfo, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stat", varargs...)
	ret0, _ := ret[0].(*fileserver.ObjectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockObjectStorageMockRecorder) Stat(ctx, key any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockObjectStorage)(nil).Stat), varargs...)
}

// Upload mocks base method.
func (m *MockObjectStorage) Upload(ctx context.Context, key string, r io.Reader, opts ...fileserver.UploadOpt) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key, r}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockObjectStorageMockRecorder) Upload(ctx, key, r any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key, r}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockObjectStorage)(nil).Upload), varargs...)
}
