// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	model "github.com/brcodingdev/chat-service/internal/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// ChatRoom is an autogenerated mock type for the ChatRoom type
type ChatRoom struct {
	mock.Mock
}

// Add provides a mock function with given fields: chatRoom
func (_m *ChatRoom) Add(chatRoom *model.ChatRoom) error {
	ret := _m.Called(chatRoom)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ChatRoom) error); ok {
		r0 = rf(chatRoom)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *ChatRoom) List() ([]model.ChatRoom, error) {
	ret := _m.Called()

	var r0 []model.ChatRoom
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.ChatRoom, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.ChatRoom); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ChatRoom)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChatRoom creates a new instance of ChatRoom. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatRoom(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatRoom {
	mock := &ChatRoom{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
