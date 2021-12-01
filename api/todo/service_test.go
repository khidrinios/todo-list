package todo_test

import (
	"fmt"
	"khidr/todo/api/todo"
	todoI "khidr/todo/interfaces/api/todo"
	mocks "khidr/todo/mocks/interfaces/api/todo"
	"khidr/todo/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type serviceMocks struct {
	repoMock *mocks.Repository
}

func setupServiceMocks() (*todo.Service, serviceMocks) {
	repoMock := new(mocks.Repository)
	s := todo.NewService(repoMock)
	return s, serviceMocks{repoMock}
}

func TestService_CreateTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s, m := setupServiceMocks()
		title := "Todo 1"
		from, to := todo.GetBeginningAndEndTime(time.Now())
		queryReq := todoI.QueryTodosRequest{
			Title: &title,
			From:  &from,
			To:    &to,
		}
		m.repoMock.On("QueryTodos", queryReq).Return([]models.Todo{}, nil)
		id := 1
		todoModel := models.Todo{
			Title: title,
		}
		m.repoMock.On("Create", todoModel).Return(&id, nil)
		request := todoI.CreateRequest{
			Title: title,
		}
		res, err := s.Create(request)
		assert.Equal(t, id, *res)
		assert.Nil(t, err)
	})

	t.Run("error - existing todo on the same day", func(t *testing.T) {
		s, m := setupServiceMocks()
		title := "Todo 1"
		from, to := todo.GetBeginningAndEndTime(time.Now())
		queryReq := todoI.QueryTodosRequest{
			Title: &title,
			From:  &from,
			To:    &to,
		}
		todoModel := models.Todo{
			Title: title,
		}
		m.repoMock.On("QueryTodos", queryReq).Return([]models.Todo{todoModel}, nil)
		request := todoI.CreateRequest{
			Title: title,
		}
		res, err := s.Create(request)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		_, ok := err.(todo.ValidationError)
		assert.True(t, ok)
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
