// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	time "time"

	models "github.com/GaryBrownEEngr/go_web_dev/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// TokenMaker is an autogenerated mock type for the TokenMaker type
type TokenMaker struct {
	mock.Mock
}

// Create provides a mock function with given fields: username, duration
func (_m *TokenMaker) Create(username string, duration time.Duration) (*models.Token, error) {
	ret := _m.Called(username, duration)

	var r0 *models.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Duration) (*models.Token, error)); ok {
		return rf(username, duration)
	}
	if rf, ok := ret.Get(0).(func(string, time.Duration) *models.Token); ok {
		r0 = rf(username, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string, time.Duration) error); ok {
		r1 = rf(username, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Verify provides a mock function with given fields: token
func (_m *TokenMaker) Verify(token *models.Token) (*models.Payload, error) {
	ret := _m.Called(token)

	var r0 *models.Payload
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Token) (*models.Payload, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(*models.Token) *models.Payload); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payload)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Token) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenMaker creates a new instance of TokenMaker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenMaker(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenMaker {
	mock := &TokenMaker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
