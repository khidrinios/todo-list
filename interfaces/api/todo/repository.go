package todo

import "khidr/todo/models"

type Repository interface {
	Create(todo models.Todo) (*int, error)
	GetTodoById(id int) (*models.Todo, error)
	UpdateTodo(id int, title, description *string, isDone *bool) (*models.Todo, error)
	QueryTodos(req models.QueryTodosRequest) ([]models.Todo, error)
	DeleteTodoById(id int) (*int, error)
}
