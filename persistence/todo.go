package persistence

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string
	Description *string
	IsDone      bool
}

func (p *PostgresConfig) CreateTodo(title string, description *string) (*int, error) {
	todo := Todo{
		Title:       title,
		Description: description,
		IsDone:      false,
	}
	result := p.db.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	todoId := int(todo.ID)
	return &todoId, nil
}
