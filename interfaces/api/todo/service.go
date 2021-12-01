package todo

import "khidr/todo/models"

type Service interface {
	Create(req CreateRequest) (*int, error)
	GetById(req TodoByIdRequestUri) (*models.Todo, error)
	List(req ListRequestBody) ([]models.Todo, error)
	DeleteById(req TodoByIdRequestUri) (*int, error)
	Update(reqParam TodoByIdRequestUri, reqbody UpdateTodoRequestBody) (*models.Todo, error)
}
