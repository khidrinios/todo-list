package todo

import (
	iterface "khidr/todo/interfaces/api/todo"
	"khidr/todo/models"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(d *gorm.DB) *Repository {
	return &Repository{
		db: d,
	}
}

func (p *Repository) Create(todo models.Todo) (*int, error) {
	result := p.db.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	todoId := int(todo.ID)
	return &todoId, nil
}

func (p *Repository) GetTodoById(id int) (*models.Todo, error) {
	todo := models.Todo{}
	result := p.db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (p *Repository) UpdateTodo(id int, title, description *string, isDone *bool) (*models.Todo, error) {
	todo := models.Todo{ID: uint(id)}
	updateBody := models.Todo{}
	if title != nil {
		updateBody.Title = *title
	}
	if description != nil {
		updateBody.Description = description
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

func (p *Repository) QueryTodos(req iterface.QueryTodosRequest) ([]models.Todo, error) {
	var todos []models.Todo
	queryTodo := models.Todo{}
	if req.IsDone != nil {
		queryTodo.IsDone = *req.IsDone
	}
	filter := BuildSqlFilterString(req.Title, req.Description, req.From, req.To)
	result := p.db.Model(&models.Todo{}).Where(filter).Limit(req.Limit).Offset(req.Offset).Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (p *Repository) DeleteTodoById(id int) (*int, error) {
	result := p.db.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &id, nil
}

func BuildSqlFilterString(title, description *string, from, to *time.Time) string {
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

	if from != nil {
		if title != nil || description != nil {
			sb.WriteString(" AND ")
		}
		sb.WriteString("created_at >= '")
		sb.WriteString((*from).Format(time.RFC3339))
		sb.WriteString("'")
	}

	if to != nil {
		if title != nil || description != nil || from != nil {
			sb.WriteString(" AND ")
		}
		sb.WriteString("created_at <= '")
		sb.WriteString((*to).Format(time.RFC3339))
		sb.WriteString("'")
	}

	return sb.String()
}
