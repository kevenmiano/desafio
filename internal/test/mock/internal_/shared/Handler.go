// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	events "github.com/aws/aws-lambda-go/events"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Handler provides a mock function with given fields: ctx, request
func (_m *Handler) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Handler")
	}

	var r0 events.APIGatewayProxyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, events.APIGatewayProxyRequest) events.APIGatewayProxyResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(events.APIGatewayProxyResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, events.APIGatewayProxyRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
