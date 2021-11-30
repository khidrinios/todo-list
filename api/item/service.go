package item

import (
	"khidr/todo/interfaces/api/item"
)

type Service struct {
	repo item.Repository
}

func NewService(postgres item.Repository) *Service {
	return &Service{
		repo: postgres,
	}
}

func (s *Service) AddToTodo(reqParam item.TodoIdRequestUri, reqbody item.AddItemToTodoRequestBody) (*item.ItemIdResult, error) {
	itemId, err := s.repo.AddItemToTodo(reqParam.TodoID, reqbody.Title, reqbody.Description)
	if err != nil {
		return nil, err
	}
	return &item.ItemIdResult{Id: *itemId}, nil
}

func (s *Service) Get(req item.ItemIdTodoIdRequestUri) (*item.ItemResult, error) {
	itemRes, err := s.repo.GetItem(req.TodoID, req.ItemID)
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

func (s *Service) Delete(req item.ItemIdTodoIdRequestUri) (*item.ItemIdResult, error) {
	itemId, err := s.repo.DeleteItem(req.TodoID, req.ItemID)
	if err != nil {
		return nil, err
	}
	return &item.ItemIdResult{
		Id: *itemId,
	}, nil
}
