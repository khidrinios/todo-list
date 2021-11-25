package todo

import (
	"khidr/todo/interfaces/api/todo"
	"khidr/todo/persistence"
)

type Service struct {
	postgres *persistence.PostgresConfig
}

func NewService(postgres *persistence.PostgresConfig) *Service {
	return &Service{
		postgres: postgres,
	}
}

func (s *Service) CreateTodo(req todo.CreateTodoRequestBody) (todo.CreateTodoResult, error) {
	todoId, err := s.postgres.CreateTodo(req.Title, req.Description)
	if err != nil {
		return todo.CreateTodoResult{}, err
	}
	return todo.CreateTodoResult{Id: *todoId}, nil
}
