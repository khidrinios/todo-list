package item

import "khidr/todo/models"

type Repository interface {
	AddItemToTodo(todoId int, title string, description *string) (*int, error)
	GetItem(todoId, itemId int) (*models.Item, error)
	DeleteItem(todoId, itemId int) (*int, error)
	GetItemsByTodoId(todoId int) ([]models.Item, error)
}
