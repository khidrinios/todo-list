package todo_test

import (
	i "khidr/todo/interfaces/api/todo"
	mocks "khidr/todo/mocks/interfaces/api/todo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_CreateTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := new(mocks.Service)
		request := i.CreateTodoRequestBody{Title: "Todo 1"}
		mockService.On("CreateTodo", request).Return(&i.TodoIdResult{
			Id: 1,
		}, nil)
		res, err := mockService.CreateTodo(request)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
}
