// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/services/filemanager (interfaces: FileManagerFactory,FileManager)

// Package mock_filemanager is a generated GoMock package.
package mock_filemanager

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	filemanager "github.com/rudderlabs/rudder-server/services/filemanager"
)

// MockFileManagerFactory is a mock of FileManagerFactory interface.
type MockFileManagerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFileManagerFactoryMockRecorder
}

// MockFileManagerFactoryMockRecorder is the mock recorder for MockFileManagerFactory.
type MockFileManagerFactoryMockRecorder struct {
	mock *MockFileManagerFactory
}

// NewMockFileManagerFactory creates a new mock instance.
func NewMockFileManagerFactory(ctrl *gomock.Controller) *MockFileManagerFactory {
	mock := &MockFileManagerFactory{ctrl: ctrl}
	mock.recorder = &MockFileManagerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileManagerFactory) EXPECT() *MockFileManagerFactoryMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockFileManagerFactory) New(arg0 *filemanager.SettingsT) (filemanager.FileManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0)
	ret0, _ := ret[0].(filemanager.FileManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockFileManagerFactoryMockRecorder) New(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockFileManagerFactory)(nil).New), arg0)
}

// MockFileManager is a mock of FileManager interface.
type MockFileManager struct {
	ctrl     *gomock.Controller
	recorder *MockFileManagerMockRecorder
}

// MockFileManagerMockRecorder is the mock recorder for MockFileManager.
type MockFileManagerMockRecorder struct {
	mock *MockFileManager
}

// NewMockFileManager creates a new mock instance.
func NewMockFileManager(ctrl *gomock.Controller) *MockFileManager {
	mock := &MockFileManager{ctrl: ctrl}
	mock.recorder = &MockFileManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileManager) EXPECT() *MockFileManagerMockRecorder {
	return m.recorder
}

// DeleteObjects mocks base method.
func (m *MockFileManager) DeleteObjects(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjects", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjects indicates an expected call of DeleteObjects.
func (mr *MockFileManagerMockRecorder) DeleteObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjects", reflect.TypeOf((*MockFileManager)(nil).DeleteObjects), arg0)
}

// Download mocks base method.
func (m *MockFileManager) Download(arg0 *os.File, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download.
func (mr *MockFileManagerMockRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockFileManager)(nil).Download), arg0, arg1)
}

// GetDownloadKeyFromFileLocation mocks base method.
func (m *MockFileManager) GetDownloadKeyFromFileLocation(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownloadKeyFromFileLocation", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDownloadKeyFromFileLocation indicates an expected call of GetDownloadKeyFromFileLocation.
func (mr *MockFileManagerMockRecorder) GetDownloadKeyFromFileLocation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadKeyFromFileLocation", reflect.TypeOf((*MockFileManager)(nil).GetDownloadKeyFromFileLocation), arg0)
}

// GetObjectNameFromLocation mocks base method.
func (m *MockFileManager) GetObjectNameFromLocation(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectNameFromLocation", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectNameFromLocation indicates an expected call of GetObjectNameFromLocation.
func (mr *MockFileManagerMockRecorder) GetObjectNameFromLocation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectNameFromLocation", reflect.TypeOf((*MockFileManager)(nil).GetObjectNameFromLocation), arg0)
}

// Upload mocks base method.
func (m *MockFileManager) Upload(arg0 *os.File, arg1 ...string) (filemanager.UploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(filemanager.UploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockFileManagerMockRecorder) Upload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockFileManager)(nil).Upload), varargs...)
}
