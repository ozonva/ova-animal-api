// Code generated by MockGen. DO NOT EDIT.
// Source: animal_flusher.go

// Package animal is a generated GoMock package.
package animal

import (
	domain "ova-animal-api/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFlusher is a mock of Flusher interface.
type MockFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockFlusherMockRecorder
}

// MockFlusherMockRecorder is the mock recorder for MockFlusher.
type MockFlusherMockRecorder struct {
	mock *MockFlusher
}

// NewMockFlusher creates a new mock instance.
func NewMockFlusher(ctrl *gomock.Controller) *MockFlusher {
	mock := &MockFlusher{ctrl: ctrl}
	mock.recorder = &MockFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlusher) EXPECT() *MockFlusherMockRecorder {
	return m.recorder
}

// Flush mocks base method.
func (m *MockFlusher) Flush(entities []domain.Animal) []domain.Animal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", entities)
	ret0, _ := ret[0].([]domain.Animal)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockFlusherMockRecorder) Flush(entities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockFlusher)(nil).Flush), entities)
}
