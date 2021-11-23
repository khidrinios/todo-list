package persistence

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

func (p *PostgresConfig) GetTodoById(id int) (*Todo, error) {
	todo := Todo{}
	result := p.db.Where("deleted_at is null").First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (p *PostgresConfig) UpdateTodo(id int, title, description *string, isDone *bool) (*Todo, error) {
	todo := Todo{Model: gorm.Model{ID: uint(id)}}
	updateBody := Todo{}
	if title != nil {
		updateBody.Title = *title
	}
	if description != nil {
		updateBody.Description = *&description
	}
	if isDone != nil {
		updateBody.IsDone = *isDone
	}
	result := p.db.Model(&todo).Clauses(clause.Returning{}).Updates(updateBody)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (p *PostgresConfig) QueryTodos(title, description *string, isDone *bool, offset, limit int) ([]Todo, error) {
	var todos []Todo
	queryTodo := Todo{}
	if isDone != nil {
		queryTodo.IsDone = *isDone
	}
	var sb strings.Builder
	if title != nil {
		sb.WriteString("title LIKE '%")
		sb.WriteString(*title)
		sb.WriteString("%'")
	}

	if description != nil {
		if title != nil {
			sb.WriteString(" AND ")
		}
		sb.WriteString("description LIKE '%")
		sb.WriteString(*description)
		sb.WriteString("%'")
	}

	if title != nil || description != nil {
		sb.WriteString(" AND ")
	}
	sb.WriteString("deleted_at is null")

	result := p.db.Model(&Todo{}).Where(sb.String()).Limit(limit).Offset(offset).Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}
