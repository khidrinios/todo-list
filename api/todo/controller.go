package todo

import (
	"errors"
	"khidr/todo/interfaces/api/todo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (ctrl Controller) GetTodoById(c *gin.Context) {
	req, err := ctrl.handler.GetTodoById(c)
	if err != nil {
		handleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.GetTodoById(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleResponseError(c, http.StatusNotFound, err)
			return
		}
		handleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) QueryTodos(c *gin.Context) {
	req, err := ctrl.handler.QueryTodos(c)
	if err != nil {
		handleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.QueryTodos(req)
	if err != nil {
		handleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) DeleteTodoById(c *gin.Context) {
	req, err := ctrl.handler.DeleteTodoById(c)
	if err != nil {
		handleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.DeleteTodoById(req)
	if err != nil {
		handleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) UpdateTodo(c *gin.Context) {
	reqParam, reqBody, err := ctrl.handler.UpdateTodo(c)
	if err != nil {
		handleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.UpdateTodo(*reqParam, *reqBody)
	if err != nil {
		handleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusOK, res)
}

func handleResponse(c *gin.Context, statusCode int, response interface{}) {
	c.JSON(statusCode, response)
}

func handleResponseError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
