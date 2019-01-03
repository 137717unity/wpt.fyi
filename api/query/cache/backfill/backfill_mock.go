// Code generated by MockGen. DO NOT EDIT.
// Source: api/query/cache/backfill/backfill.go

// Package backfill is a generated GoMock package.
package backfill

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
)

// MockRunFetcher is a mock of RunFetcher interface
type MockRunFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockRunFetcherMockRecorder
}

// MockRunFetcherMockRecorder is the mock recorder for MockRunFetcher
type MockRunFetcherMockRecorder struct {
	mock *MockRunFetcher
}

// NewMockRunFetcher creates a new mock instance
func NewMockRunFetcher(ctrl *gomock.Controller) *MockRunFetcher {
	mock := &MockRunFetcher{ctrl: ctrl}
	mock.recorder = &MockRunFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunFetcher) EXPECT() *MockRunFetcherMockRecorder {
	return m.recorder
}

// FetchRuns mocks base method
func (m *MockRunFetcher) FetchRuns(limit int) (shared.TestRunsByProduct, error) {
	ret := m.ctrl.Call(m, "FetchRuns", limit)
	ret0, _ := ret[0].(shared.TestRunsByProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRuns indicates an expected call of FetchRuns
func (mr *MockRunFetcherMockRecorder) FetchRuns(limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRuns", reflect.TypeOf((*MockRunFetcher)(nil).FetchRuns), limit)
}
