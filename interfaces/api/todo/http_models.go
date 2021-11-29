package todo

import "time"

type CreateTodoRequestBody struct {
	Description *string `json:"description,omitempty"`
	Title       string  `json:"title" binding:"required"`
}

type TodoIdResult struct {
	Id int `json:"id" binding:"required"`
}

type TodoByIdRequestUri struct {
	ID int `uri:"todo_id" binding:"required,min=1"`
}

type Item struct {
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	Description *string   `json:"description,omitempty"`
	Id          int       `json:"id" binding:"required"`
	IsDone      bool      `json:"is_done" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
}

type TodoResult struct {
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	Description *string   `json:"description,omitempty"`
	Id          int       `json:"id" binding:"required"`
	IsDone      bool      `json:"is_done" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
	Items       []Item    `json:"items"`
}

type QueryTodosRequestBody struct {
	Description *string `json:"description,omitempty"`
	IsDone      *bool   `json:"is_done,omitempty"`
	Limit       int     `json:"limit" binding:"required,min=1,max=1000"`
	Page        int     `json:"page"  binding:"required,min=1"`
	Title       *string `json:"title,omitempty"`
}

type UpdateTodoRequestBody struct {
	Description *string `json:"description,omitempty"`
	IsDone      *bool   `json:"is_done,omitempty"`
	Title       *string `json:"title,omitempty"`
}

type UpdateTodoResult struct {
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	Description *string   `json:"description,omitempty"`
	Id          int       `json:"id" binding:"required"`
	IsDone      bool      `json:"is_done" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
}
