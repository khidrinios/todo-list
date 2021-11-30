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

func (ctrl Controller) AddToTodo(c *gin.Context) {
	reqParam, reqBody, err := ctrl.handler.AddToTodo(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}
	res, err := ctrl.service.AddToTodo(*reqParam, *reqBody)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusCreated, res)
}

func (ctrl Controller) Get(c *gin.Context) {
	req, err := ctrl.handler.Get(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}
	res, err := ctrl.service.Get(*req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) Delete(c *gin.Context) {
	req, err := ctrl.handler.Delete(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}
	res, err := ctrl.service.Delete(*req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}
