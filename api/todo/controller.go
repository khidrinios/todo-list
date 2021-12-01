package todo

import (
	"errors"
	"khidr/todo/api"
	itemI "khidr/todo/interfaces/api/item"
	todoI "khidr/todo/interfaces/api/todo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	handler Handler
	todoSvc todoI.Service
	itemSvc itemI.Service
}

func NewController(todoSvc todoI.Service, itemSvc itemI.Service, handler Handler) Controller {
	return Controller{
		todoSvc: todoSvc,
		itemSvc: itemSvc,
		handler: handler,
	}
}

func (ctrl Controller) Create(c *gin.Context) {
	req, err := ctrl.handler.Create(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}
	todoId, err := ctrl.todoSvc.Create(req)
	if err != nil {
		if _, ok := err.(ValidationError); ok {
			api.HandleResponseError(c, http.StatusUnprocessableEntity, err)
			return
		}
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	res := todoI.TodoIdResult{
		Id: *todoId,
	}
	api.HandleResponse(c, http.StatusCreated, res)
}

func (ctrl Controller) GetById(c *gin.Context) {
	req, err := ctrl.handler.GetById(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}
	res, err := ctrl.todoSvc.GetById(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.HandleResponseError(c, http.StatusNotFound, err)
			return
		}
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	items, err := ctrl.itemSvc.GetItemsByTodoId(req.ID)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	itemsRes := make([]todoI.Item, len(items))
	for i := range items {
		itemRes := todoI.Item{
			CreatedAt:   items[i].CreatedAt,
			Description: items[i].Description,
			Id:          int(items[i].ID),
			IsDone:      items[i].IsDone,
			Title:       items[i].Title,
			UpdatedAt:   items[i].UpdatedAt,
		}
		itemsRes = append(itemsRes, itemRes)
	}
	todoRes := todoI.TodoResult{
		CreatedAt:   res.CreatedAt,
		Description: res.Description,
		Id:          req.ID,
		IsDone:      res.IsDone,
		Title:       res.Title,
		UpdatedAt:   res.UpdatedAt,
		Items:       itemsRes,
	}
	api.HandleResponse(c, http.StatusOK, todoRes)
}

func (ctrl Controller) List(c *gin.Context) {
	req, err := ctrl.handler.List(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	res, err := ctrl.todoSvc.List(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	todosRes := make([]todoI.TodoResult, len(res))
	for i := range res {
		todoRes := todoI.TodoResult{
			CreatedAt:   res[i].CreatedAt,
			Description: res[i].Description,
			Id:          int(res[i].ID),
			IsDone:      res[i].IsDone,
			Title:       res[i].Title,
			UpdatedAt:   res[i].UpdatedAt,
		}
		todosRes = append(todosRes, todoRes)
	}
	api.HandleResponse(c, http.StatusOK, todosRes)
}

func (ctrl Controller) DeleteById(c *gin.Context) {
	req, err := ctrl.handler.DeleteById(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	todoId, err := ctrl.todoSvc.DeleteById(req)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	res := todoI.TodoIdResult{
		Id: *todoId,
	}
	api.HandleResponse(c, http.StatusOK, res)
}

func (ctrl Controller) Update(c *gin.Context) {
	reqParam, reqBody, err := ctrl.handler.Update(c)
	if err != nil {
		api.HandleResponseError(c, http.StatusBadRequest, err)
		return
	}

	todo, err := ctrl.todoSvc.Update(*reqParam, *reqBody)
	if err != nil {
		api.HandleResponseError(c, http.StatusInternalServerError, err)
		return
	}
	res := todoI.UpdateTodoResult{
		CreatedAt:   todo.CreatedAt,
		Description: todo.Description,
		Id:          int(todo.ID),
		IsDone:      todo.IsDone,
		Title:       todo.Title,
		UpdatedAt:   todo.UpdatedAt,
	}
	api.HandleResponse(c, http.StatusOK, res)
}
