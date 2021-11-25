package todo

type CreateTodoRequestBody struct {
	Description *string `json:"description,omitempty"`
	Title       string  `json:"title" binding:"required"`
}

type CreateTodoResult struct {
	Id int `json:"id" binding:"required"`
}
