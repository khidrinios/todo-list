package item

import (
	"khidr/todo/api"
	"khidr/todo/interfaces/api/item"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	handler Handler
	service item.Service
}

func NewController(service item.Service, handler Handler) Controller {
	return Controller{
		service: service,
		handler: handler,
	}
}

func (ctrl Controller) AddItemToTodo(c *gin.Context) {
	reqParam, reqBody, err := ctrl.handler.AddItemToTodo(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}
	res, err := ctrl.service.AddItemToTodo(*reqParam, *reqBody)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusCreated, res)
}

func (ctrl Controller) GetItemByIdAndTodoId(c *gin.Context) {
	req, err := ctrl.handler.GetItemByIdAndTodoId(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}
	res, err := ctrl.service.GetItemByIdAndTodoId(*req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}
