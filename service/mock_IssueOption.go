// Code generated by mockery v2.41.0. DO NOT EDIT.

package service

import mock "github.com/stretchr/testify/mock"

// MockIssueOption is an autogenerated mock type for the IssueOption type
type MockIssueOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *MockIssueOption) Execute(_a0 *issueOption) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*issueOption) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockIssueOption creates a new instance of MockIssueOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIssueOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIssueOption {
	mock := &MockIssueOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
