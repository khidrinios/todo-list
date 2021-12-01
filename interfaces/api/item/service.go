package item

import "khidr/todo/models"

type Service interface {
	AddToTodo(reqParam TodoIdRequestUri, reqbody AddItemToTodoRequestBody) (*ItemIdResult, error)
	Get(req ItemIdTodoIdRequestUri) (*ItemResult, error)
	Delete(req ItemIdTodoIdRequestUri) (*ItemIdResult, error)
	GetItemsByTodoId(todoId int) ([]models.Item, error)
}
