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

func (ctrl Controller) Create(c *gin.Context) {
	req, err := ctrl.handler.Create(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.Create(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusCreated, res)
}

func (ctrl Controller) GetById(c *gin.Context) {
	req, err := ctrl.handler.GetById(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.GetById(req)
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

func (ctrl Controller) List(c *gin.Context) {
	req, err := ctrl.handler.List(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.List(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) DeleteById(c *gin.Context) {
	req, err := ctrl.handler.DeleteById(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.DeleteById(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) Update(c *gin.Context) {
	reqParam, reqBody, err := ctrl.handler.Update(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.service.Update(*reqParam, *reqBody)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	api.HandleResponse(c, http.StatusOK, res)
}
