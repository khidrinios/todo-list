package item

type Service interface {
	AddItemToTodo(reqParam TodoIdRequestUri, reqbody AddItemToTodoRequestBody) (*ItemIdResult, error)
	GetItem(req ItemIdTodoIdRequestUri) (*ItemResult, error)
	DeleteItem(req ItemIdTodoIdRequestUri) (*ItemIdResult, error)
}
