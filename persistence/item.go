package persistence

import (
	"database/sql"
	"time"
)

type Item struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"uniqueIndex:item_idx_title_createdat"`
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
	Title       string       `gorm:"uniqueIndex:item_idx_title_createdat"`
	Description *string
	IsDone      bool
	TodoID      int `gorm:"primarykey"`
	Todo        Todo
}

func (p *PostgresConfig) AddItemToTodo(todoId int, title string, description *string) (*int, error) {
	item := Item{
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

func (p *PostgresConfig) GetItem(todoId, itemId int) (*Item, error) {
	item := Item{}
	result := p.db.First(&item, Item{ID: uint(itemId), TodoID: todoId})
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (p *PostgresConfig) DeleteItem(todoId, itemId int) (*int, error) {
	result := p.db.Delete(&Item{}, Item{ID: uint(itemId), TodoID: todoId})
	if result.Error != nil {
		return nil, result.Error
	}
	return &itemId, nil
}
