package item

import (
	"khidr/todo/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(d *gorm.DB) *Repository {
	return &Repository{
		db: d,
	}
}

func (p *Repository) AddItemToTodo(todoId int, title string, description *string) (*int, error) {
	item := models.Item{
		Title:       title,
		Description: description,
		TodoID:      todoId,
	}
	result := p.db.Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	itemId := int(item.ID)
	return &itemId, nil
}

func (p *Repository) GetItem(todoId, itemId int) (*models.Item, error) {
	item := models.Item{}
	result := p.db.First(&item, models.Item{ID: uint(itemId), TodoID: todoId})
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (p *Repository) DeleteItem(todoId, itemId int) (*int, error) {
	result := p.db.Delete(&models.Item{}, models.Item{ID: uint(itemId), TodoID: todoId})
	if result.Error != nil {
		return nil, result.Error
	}
	return &itemId, nil
}

func (p *Repository) GetItemsByTodoId(todoId int) ([]models.Item, error) {
	var items []models.Item
	result := p.db.Find(&items, models.Item{TodoID: todoId})
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}
