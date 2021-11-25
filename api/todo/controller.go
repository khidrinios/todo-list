package todo

import (
	"khidr/todo/interfaces/api/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	handler Handler
	service todo.Service
}

func NewController(service todo.Service, handler Handler) Controller {
	return Controller{
		service: service,
		handler: handler,
	}
}

func (ctrl Controller) CreateTodo(c *gin.Context) {
	req, err := ctrl.handler.CreateTodo(c)
	if err != nil {
		handleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.CreateTodo(req)
	if err != nil {
		handleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusCreated, res)
}

func handleResponse(c *gin.Context, statusCode int, response interface{}) {
	c.JSON(statusCode, response)
}

func handleResponseError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
