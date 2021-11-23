package service

import "khidr/todo/persistence"

type Service struct {
	postgres *persistence.PostgresConfig
}

func New(postgres *persistence.PostgresConfig) *Service {
	return &Service{
		postgres: postgres,
	}
}

func (s *Service) CreateTodo(title string, description *string) (*int, error) {
	todoId, err := s.postgres.CreateTodo(title, description)
	if err != nil {
		return nil, err
	}
	return todoId, nil
}
