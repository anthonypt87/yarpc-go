// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/api/transport (interfaces: Router,RouteTable)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package transporttest is a generated GoMock package.
package transporttest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/api/transport"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouter) Choose(arg0 context.Context, arg1 *transport.Request) (transport.HandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(transport.HandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouterMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouter)(nil).Choose), arg0, arg1)
}

// Procedures mocks base method
func (m *MockRouter) Procedures() []transport.Procedure {
	ret := m.ctrl.Call(m, "Procedures")
	ret0, _ := ret[0].([]transport.Procedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouterMockRecorder) Procedures() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouter)(nil).Procedures))
}

// MockRouteTable is a mock of RouteTable interface
type MockRouteTable struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTableMockRecorder
}

// MockRouteTableMockRecorder is the mock recorder for MockRouteTable
type MockRouteTableMockRecorder struct {
	mock *MockRouteTable
}

// NewMockRouteTable creates a new mock instance
func NewMockRouteTable(ctrl *gomock.Controller) *MockRouteTable {
	mock := &MockRouteTable{ctrl: ctrl}
	mock.recorder = &MockRouteTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteTable) EXPECT() *MockRouteTableMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouteTable) Choose(arg0 context.Context, arg1 *transport.Request) (transport.HandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(transport.HandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouteTableMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouteTable)(nil).Choose), arg0, arg1)
}

// Procedures mocks base method
func (m *MockRouteTable) Procedures() []transport.Procedure {
	ret := m.ctrl.Call(m, "Procedures")
	ret0, _ := ret[0].([]transport.Procedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouteTableMockRecorder) Procedures() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouteTable)(nil).Procedures))
}

// Register mocks base method
func (m *MockRouteTable) Register(arg0 []transport.Procedure) {
	m.ctrl.Call(m, "Register", arg0)
}

// Register indicates an expected call of Register
func (mr *MockRouteTableMockRecorder) Register(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRouteTable)(nil).Register), arg0)
}