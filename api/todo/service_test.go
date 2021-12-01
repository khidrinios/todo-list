package todo_test

import (
	"fmt"
	"khidr/todo/api/todo"
	i "khidr/todo/interfaces/api/todo"
	mocks "khidr/todo/mocks/interfaces/api/todo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestService_CreateTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := new(mocks.Service)
		request := i.CreateRequest{Title: "Todo 1"}
		mockService.On("CreateTodo", request).Return(&i.TodoIdResult{
			Id: 1,
		}, nil)
		res, err := mockService.CreateTodo(request)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
}

func Test_BeginningAndEndTime(t *testing.T) {
	now := time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
	fmt.Println(now)
	fromActual, toActual := todo.GetBeginningAndEndTime(now)
	fromExpected := time.Date(2021, time.Month(2), 21, 0, 0, 0, 0, time.UTC)
	toExpected := time.Date(2021, time.Month(2), 22, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, fromExpected, fromActual)
	assert.Equal(t, toExpected, toActual)
}
