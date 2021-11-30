package todo

type Service interface {
	Create(req CreateRequest) (*TodoIdResult, error)
	GetTodoById(req TodoByIdRequestUri) (*TodoResult, error)
	QueryTodos(req QueryTodosRequestBody) ([]TodoResult, error)
	DeleteTodoById(req TodoByIdRequestUri) (*TodoIdResult, error)
	UpdateTodo(reqParam TodoByIdRequestUri, reqbody UpdateTodoRequestBody) (*UpdateTodoResult, error)
}
