package todo

import (
	"khidr/todo/interfaces/api/todo"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Create(c *gin.Context) (todo.CreateRequest, error) {
	request := new(todo.CreateRequest)
	err := c.ShouldBindJSON(request)
	return *request, err
}

func (h Handler) GetTodoById(c *gin.Context) (todo.TodoByIdRequestUri, error) {
	request := new(todo.TodoByIdRequestUri)
	err := c.ShouldBindUri(request)
	return *request, err
}

func (h Handler) QueryTodos(c *gin.Context) (todo.QueryTodosRequestBody, error) {
	request := new(todo.QueryTodosRequestBody)
	err := c.ShouldBindJSON(request)
	return *request, err
}

func (h Handler) DeleteTodoById(c *gin.Context) (todo.TodoByIdRequestUri, error) {
	request := new(todo.TodoByIdRequestUri)
	err := c.ShouldBindUri(request)
	return *request, err
}

func (h Handler) UpdateTodo(c *gin.Context) (*todo.TodoByIdRequestUri, *todo.UpdateTodoRequestBody, error) {
	requestParam := new(todo.TodoByIdRequestUri)
	err := c.ShouldBindUri(requestParam)
	if err != nil {
		return nil, nil, err
	}
	requestBody := new(todo.UpdateTodoRequestBody)
	err = c.ShouldBindJSON(requestBody)
	return requestParam, requestBody, err
}
