package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"uniqueIndex:todo_idx_title_createdat"`
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
	Title       string       `gorm:"uniqueIndex:todo_idx_title_createdat"`
	Description *string
	IsDone      bool
}

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
