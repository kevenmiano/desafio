// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	events "github.com/aws/aws-lambda-go/events"
	mock "github.com/stretchr/testify/mock"
)

// Multipart is an autogenerated mock type for the Multipart type
type Multipart[T interface{}] struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, request, file
func (_m *Multipart[T]) Execute(ctx context.Context, request events.APIGatewayProxyRequest, file T) (events.APIGatewayProxyResponse, error) {
	ret := _m.Called(ctx, request, file)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 events.APIGatewayProxyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, events.APIGatewayProxyRequest, T) (events.APIGatewayProxyResponse, error)); ok {
		return rf(ctx, request, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, events.APIGatewayProxyRequest, T) events.APIGatewayProxyResponse); ok {
		r0 = rf(ctx, request, file)
	} else {
		r0 = ret.Get(0).(events.APIGatewayProxyResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, events.APIGatewayProxyRequest, T) error); ok {
		r1 = rf(ctx, request, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMultipart creates a new instance of Multipart. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMultipart[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Multipart[T] {
	mock := &Multipart[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
