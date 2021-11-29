// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	todo "khidr/todo/interfaces/api/todo"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: req
func (_m *Service) CreateTodo(req todo.CreateTodoRequestBody) (*todo.TodoIdResult, error) {
	ret := _m.Called(req)

	var r0 *todo.TodoIdResult
	if rf, ok := ret.Get(0).(func(todo.CreateTodoRequestBody) *todo.TodoIdResult); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.TodoIdResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(todo.CreateTodoRequestBody) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTodoById provides a mock function with given fields: req
func (_m *Service) DeleteTodoById(req todo.TodoByIdRequestUri) (*todo.TodoIdResult, error) {
	ret := _m.Called(req)

	var r0 *todo.TodoIdResult
	if rf, ok := ret.Get(0).(func(todo.TodoByIdRequestUri) *todo.TodoIdResult); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.TodoIdResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(todo.TodoByIdRequestUri) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoById provides a mock function with given fields: req
func (_m *Service) GetTodoById(req todo.TodoByIdRequestUri) (*todo.TodoResult, error) {
	ret := _m.Called(req)

	var r0 *todo.TodoResult
	if rf, ok := ret.Get(0).(func(todo.TodoByIdRequestUri) *todo.TodoResult); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.TodoResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(todo.TodoByIdRequestUri) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryTodos provides a mock function with given fields: req
func (_m *Service) QueryTodos(req todo.QueryTodosRequestBody) ([]todo.TodoResult, error) {
	ret := _m.Called(req)

	var r0 []todo.TodoResult
	if rf, ok := ret.Get(0).(func(todo.QueryTodosRequestBody) []todo.TodoResult); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]todo.TodoResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(todo.QueryTodosRequestBody) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodo provides a mock function with given fields: reqParam, reqbody
func (_m *Service) UpdateTodo(reqParam todo.TodoByIdRequestUri, reqbody todo.UpdateTodoRequestBody) (*todo.UpdateTodoResult, error) {
	ret := _m.Called(reqParam, reqbody)

	var r0 *todo.UpdateTodoResult
	if rf, ok := ret.Get(0).(func(todo.TodoByIdRequestUri, todo.UpdateTodoRequestBody) *todo.UpdateTodoResult); ok {
		r0 = rf(reqParam, reqbody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.UpdateTodoResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(todo.TodoByIdRequestUri, todo.UpdateTodoRequestBody) error); ok {
		r1 = rf(reqParam, reqbody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
