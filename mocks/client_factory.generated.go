// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	client "github.com/0xPolygon/cdk-data-availability/client"
	mock "github.com/stretchr/testify/mock"
)

// IClientFactory is an autogenerated mock type for the IClientFactory type
type IClientFactory struct {
	mock.Mock
}

// New provides a mock function with given fields: url
func (_m *IClientFactory) New(url string) client.IClient {
	ret := _m.Called(url)

	var r0 client.IClient
	if rf, ok := ret.Get(0).(func(string) client.IClient); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.IClient)
		}
	}

	return r0
}

type mockConstructorTestingTNewIClientFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewIClientFactory creates a new instance of IClientFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIClientFactory(t mockConstructorTestingTNewIClientFactory) *IClientFactory {
	mock := &IClientFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}