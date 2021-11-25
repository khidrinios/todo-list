package todo

import (
	"khidr/todo/interfaces/api/todo"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) CreateTodo(c *gin.Context) (todo.CreateTodoRequestBody, error) {
	request := new(todo.CreateTodoRequestBody)
	err := c.ShouldBindJSON(request)
	return *request, err
}
