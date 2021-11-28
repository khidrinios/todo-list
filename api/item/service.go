package item

import (
	"khidr/todo/interfaces/api/item"
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

func (s *Service) AddItemToTodo(reqParam item.TodoIdRequestUri, reqbody item.AddItemToTodoRequestBody) (*item.AddItemToTodoResult, error) {
	itemId, err := s.postgres.AddItemToTodo(reqParam.TodoID, reqbody.Title, reqbody.Description)
	if err != nil {
		return nil, err
	}
	return &item.AddItemToTodoResult{Id: *itemId}, nil
}

func (s *Service) GetItemByIdAndTodoId(req item.ItemIdTodoIdRequestUri) (*item.ItemResult, error) {
	itemRes, err := s.postgres.GetItemByIdAndTodoId(req.TodoID, req.ItemID)
	if err != nil {
		return nil, err
	}
	return &item.ItemResult{
		CreatedAt:   itemRes.CreatedAt,
		Description: itemRes.Description,
		Id:          int(itemRes.ID),
		IsDone:      itemRes.IsDone,
		Title:       itemRes.Title,
		UpdatedAt:   itemRes.UpdatedAt,
	}, nil
}
