// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/services/transformer (interfaces: FeaturesService)

// Package mock_features is a generated GoMock package.
package mock_features

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeaturesService is a mock of FeaturesService interface.
type MockFeaturesService struct {
	ctrl     *gomock.Controller
	recorder *MockFeaturesServiceMockRecorder
}

// MockFeaturesServiceMockRecorder is the mock recorder for MockFeaturesService.
type MockFeaturesServiceMockRecorder struct {
	mock *MockFeaturesService
}

// NewMockFeaturesService creates a new mock instance.
func NewMockFeaturesService(ctrl *gomock.Controller) *MockFeaturesService {
	mock := &MockFeaturesService{ctrl: ctrl}
	mock.recorder = &MockFeaturesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeaturesService) EXPECT() *MockFeaturesServiceMockRecorder {
	return m.recorder
}

// Regulations mocks base method.
func (m *MockFeaturesService) Regulations() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Regulations")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Regulations indicates an expected call of Regulations.
func (mr *MockFeaturesServiceMockRecorder) Regulations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Regulations", reflect.TypeOf((*MockFeaturesService)(nil).Regulations))
}

// RouterTransform mocks base method.
func (m *MockFeaturesService) RouterTransform(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouterTransform", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RouterTransform indicates an expected call of RouterTransform.
func (mr *MockFeaturesServiceMockRecorder) RouterTransform(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouterTransform", reflect.TypeOf((*MockFeaturesService)(nil).RouterTransform), arg0)
}

// SourceTransformerVersion mocks base method.
func (m *MockFeaturesService) SourceTransformerVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourceTransformerVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// SourceTransformerVersion indicates an expected call of SourceTransformerVersion.
func (mr *MockFeaturesServiceMockRecorder) SourceTransformerVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourceTransformerVersion", reflect.TypeOf((*MockFeaturesService)(nil).SourceTransformerVersion))
}

// TransformerProxyVersion mocks base method.
func (m *MockFeaturesService) TransformerProxyVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransformerProxyVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// TransformerProxyVersion indicates an expected call of TransformerProxyVersion.
func (mr *MockFeaturesServiceMockRecorder) TransformerProxyVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransformerProxyVersion", reflect.TypeOf((*MockFeaturesService)(nil).TransformerProxyVersion))
}

// Wait mocks base method.
func (m *MockFeaturesService) Wait() chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(chan struct{})
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockFeaturesServiceMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockFeaturesService)(nil).Wait))
}