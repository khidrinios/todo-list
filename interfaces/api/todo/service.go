package todo

type Service interface {
	CreateTodo(req CreateTodoRequestBody) (*CreateTodoResult, error)
	GetTodoById(req GetTodoByIdRequestUri) (*TodoResult, error)
	QueryTodos(req QueryTodosRequestBody) ([]TodoResult, error)
	// DeleteTodoById(w http.ResponseWriter, r *http.Request, id int)
	// UpdateTodo(w http.ResponseWriter, r *http.Request, id int)
}
