// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	events "github.com/aws/aws-lambda-go/events"
	domain "github.com/kevenmiano/v3/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// AdapterInterface is an autogenerated mock type for the AdapterInterface type
type AdapterInterface struct {
	mock.Mock
}

// Read provides a mock function with given fields: request
func (_m *AdapterInterface) Read(request *events.APIGatewayProxyRequest) (*domain.File, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 *domain.File
	var r1 error
	if rf, ok := ret.Get(0).(func(*events.APIGatewayProxyRequest) (*domain.File, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*events.APIGatewayProxyRequest) *domain.File); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.File)
		}
	}

	if rf, ok := ret.Get(1).(func(*events.APIGatewayProxyRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAdapterInterface creates a new instance of AdapterInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdapterInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdapterInterface {
	mock := &AdapterInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
