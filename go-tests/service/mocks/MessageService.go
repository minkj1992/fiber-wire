// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MessageService is an autogenerated mock type for the MessageService type
type MessageService struct {
	mock.Mock
}

// SendChargeNotification provides a mock function with given fields: _a0
func (_m *MessageService) SendChargeNotification(_a0 int) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewMessageService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageService creates a new instance of MessageService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageService(t mockConstructorTestingTNewMessageService) *MessageService {
	mock := &MessageService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
