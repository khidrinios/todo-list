package todo

import "khidr/todo/models"

type Service interface {
	Create(req CreateRequest) (*TodoIdResult, error)
	GetById(req TodoByIdRequestUri) (*models.Todo, error)
	List(req QueryTodosRequestBody) ([]models.Todo, error)
	DeleteById(req TodoByIdRequestUri) (*TodoIdResult, error)
	Update(reqParam TodoByIdRequestUri, reqbody UpdateTodoRequestBody) (*UpdateTodoResult, error)
}
