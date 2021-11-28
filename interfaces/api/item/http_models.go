package item

import "time"

type TodoIdRequestUri struct {
	TodoID int `uri:"todo_id" binding:"required,min=1"`
}

type AddItemToTodoRequestBody struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description,omitempty"`
}

type AddItemToTodoResult struct {
	Id int `json:"id" binding:"required"`
}

type ItemIdTodoIdRequestUri struct {
	ItemID int `uri:"item_id" binding:"required,min=1"`
	TodoID int `uri:"todo_id" binding:"required,min=1"`
}

type ItemResult struct {
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	Description *string   `json:"description,omitempty"`
	Id          int       `json:"id" binding:"required"`
	IsDone      bool      `json:"is_done" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
}
