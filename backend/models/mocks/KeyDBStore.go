// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// KeyDBStore is an autogenerated mock type for the KeyDBStore type
type KeyDBStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, key
func (_m *KeyDBStore) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key, out
func (_m *KeyDBStore) Get(ctx context.Context, key string, out interface{}) error {
	ret := _m.Called(ctx, key, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: ctx, in
func (_m *KeyDBStore) Put(ctx context.Context, in interface{}) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKeyDBStore creates a new instance of KeyDBStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKeyDBStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *KeyDBStore {
	mock := &KeyDBStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
