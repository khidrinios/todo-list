package item

import (
	"khidr/todo/interfaces/api/item"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) AddToTodo(c *gin.Context) (*item.TodoIdRequestUri, *item.AddItemToTodoRequestBody, error) {
	requestParam := new(item.TodoIdRequestUri)
	err := c.ShouldBindUri(requestParam)
	if err != nil {
		return nil, nil, err
	}
	requestBody := new(item.AddItemToTodoRequestBody)
	err = c.ShouldBindJSON(requestBody)
	return requestParam, requestBody, err
}

func (h Handler) Get(c *gin.Context) (*item.ItemIdTodoIdRequestUri, error) {
	requestParam := new(item.ItemIdTodoIdRequestUri)
	err := c.ShouldBindUri(requestParam)
	return requestParam, err
}

func (h Handler) Delete(c *gin.Context) (*item.ItemIdTodoIdRequestUri, error) {
	requestParam := new(item.ItemIdTodoIdRequestUri)
	err := c.ShouldBindUri(requestParam)
	return requestParam, err
}
