// Code generated by mockery v2.41.0. DO NOT EDIT.

package store

import (
	context "context"

	models "github.com/A-pen-app/kickstart/models"
	mock "github.com/stretchr/testify/mock"
)

// MockOrder is an autogenerated mock type for the Order type
type MockOrder struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, orderID
func (_m *MockOrder) Delete(ctx context.Context, orderID string) error {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, orderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLiveOrders provides a mock function with given fields: ctx, action
func (_m *MockOrder) GetLiveOrders(ctx context.Context, action models.OrderAction) ([]*models.Order, error) {
	ret := _m.Called(ctx, action)

	if len(ret) == 0 {
		panic("no return value specified for GetLiveOrders")
	}

	var r0 []*models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderAction) ([]*models.Order, error)); ok {
		return rf(ctx, action)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderAction) []*models.Order); ok {
		r0 = rf(ctx, action)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.OrderAction) error); ok {
		r1 = rf(ctx, action)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Make provides a mock function with given fields: ctx, action, price, quantity
func (_m *MockOrder) Make(ctx context.Context, action models.OrderAction, price int, quantity int) error {
	ret := _m.Called(ctx, action, price, quantity)

	if len(ret) == 0 {
		panic("no return value specified for Make")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderAction, int, int) error); ok {
		r0 = rf(ctx, action, price, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Take provides a mock function with given fields: ctx, action, quantity
func (_m *MockOrder) Take(ctx context.Context, action models.OrderAction, quantity int) (int, error) {
	ret := _m.Called(ctx, action, quantity)

	if len(ret) == 0 {
		panic("no return value specified for Take")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderAction, int) (int, error)); ok {
		return rf(ctx, action, quantity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderAction, int) int); ok {
		r0 = rf(ctx, action, quantity)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.OrderAction, int) error); ok {
		r1 = rf(ctx, action, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
