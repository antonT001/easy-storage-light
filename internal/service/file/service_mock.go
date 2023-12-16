// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package fileservice is a generated GoMock package.
package fileservice

import (
	io "io"
	reflect "reflect"

	models "github.com/antonT001/easy-storage-light/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// UploadChunk mocks base method.
func (m *MockService) UploadChunk(upload models.UploadChunk, body io.ReadCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadChunk", upload, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadChunk indicates an expected call of UploadChunk.
func (mr *MockServiceMockRecorder) UploadChunk(upload, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadChunk", reflect.TypeOf((*MockService)(nil).UploadChunk), upload, body)
}
