package todo

import (
	"errors"
	"khidr/todo/api"
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
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.CreateTodo(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusCreated, res)
}

func (ctrl Controller) GetTodoById(c *gin.Context) {
	req, err := ctrl.handler.GetTodoById(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.GetTodoById(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.HandleResponseError(c, http.StatusNotFound, err)
			return
		}
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) QueryTodos(c *gin.Context) {
	req, err := ctrl.handler.QueryTodos(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.QueryTodos(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) DeleteTodoById(c *gin.Context) {
	req, err := ctrl.handler.DeleteTodoById(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.DeleteTodoById(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) UpdateTodo(c *gin.Context) {
	reqParam, reqBody, err := ctrl.handler.UpdateTodo(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.UpdateTodo(*reqParam, *reqBody)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}
