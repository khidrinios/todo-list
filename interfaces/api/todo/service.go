package todo

type Service interface {
	CreateTodo(req CreateTodoRequestBody) (CreateTodoResult, error)
}
