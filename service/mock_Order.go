// Code generated by mockery v2.41.0. DO NOT EDIT.

package service

import (
	context "context"

	models "github.com/A-pen-app/kickstart/models"
	mock "github.com/stretchr/testify/mock"
)

// MockOrder is an autogenerated mock type for the Order type
type MockOrder struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, userID, kickstartID
func (_m *MockOrder) Get(ctx context.Context, userID string, kickstartID string) (*models.Order, error) {
	ret := _m.Called(ctx, userID, kickstartID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.Order, error)); ok {
		return rf(ctx, userID, kickstartID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Order); ok {
		r0 = rf(ctx, userID, kickstartID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, kickstartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: ctx, userID, next, count, filter
func (_m *MockOrder) GetOrders(ctx context.Context, userID string, next string, count int, filter models.OrderStatus) ([]*models.Order, string, error) {
	ret := _m.Called(ctx, userID, next, count, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 []*models.Order
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, models.OrderStatus) ([]*models.Order, string, error)); ok {
		return rf(ctx, userID, next, count, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, models.OrderStatus) []*models.Order); ok {
		r0 = rf(ctx, userID, next, count, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, models.OrderStatus) string); ok {
		r1 = rf(ctx, userID, next, count, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, int, models.OrderStatus) error); ok {
		r2 = rf(ctx, userID, next, count, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// New provides a mock function with given fields: ctx, userID, kickstartID, email
func (_m *MockOrder) New(ctx context.Context, userID string, kickstartID string, email *string) error {
	ret := _m.Called(ctx, userID, kickstartID, email)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string) error); ok {
		r0 = rf(ctx, userID, kickstartID, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockOrder creates a new instance of MockOrder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrder {
	mock := &MockOrder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
