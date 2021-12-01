// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	item "khidr/todo/interfaces/api/item"

	mock "github.com/stretchr/testify/mock"

	models "khidr/todo/models"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddToTodo provides a mock function with given fields: reqParam, reqbody
func (_m *Service) AddToTodo(reqParam item.TodoIdRequestUri, reqbody item.AddItemToTodoRequestBody) (*item.ItemIdResult, error) {
	ret := _m.Called(reqParam, reqbody)

	var r0 *item.ItemIdResult
	if rf, ok := ret.Get(0).(func(item.TodoIdRequestUri, item.AddItemToTodoRequestBody) *item.ItemIdResult); ok {
		r0 = rf(reqParam, reqbody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*item.ItemIdResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(item.TodoIdRequestUri, item.AddItemToTodoRequestBody) error); ok {
		r1 = rf(reqParam, reqbody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: req
func (_m *Service) Delete(req item.ItemIdTodoIdRequestUri) (*item.ItemIdResult, error) {
	ret := _m.Called(req)

	var r0 *item.ItemIdResult
	if rf, ok := ret.Get(0).(func(item.ItemIdTodoIdRequestUri) *item.ItemIdResult); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*item.ItemIdResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(item.ItemIdTodoIdRequestUri) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: req
func (_m *Service) Get(req item.ItemIdTodoIdRequestUri) (*item.ItemResult, error) {
	ret := _m.Called(req)

	var r0 *item.ItemResult
	if rf, ok := ret.Get(0).(func(item.ItemIdTodoIdRequestUri) *item.ItemResult); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*item.ItemResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(item.ItemIdTodoIdRequestUri) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItemsByTodoId provides a mock function with given fields: todoId
func (_m *Service) GetItemsByTodoId(todoId int) ([]models.Item, error) {
	ret := _m.Called(todoId)

	var r0 []models.Item
	if rf, ok := ret.Get(0).(func(int) []models.Item); ok {
		r0 = rf(todoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(todoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
