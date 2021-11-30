package item

type Service interface {
	AddToTodo(reqParam TodoIdRequestUri, reqbody AddItemToTodoRequestBody) (*ItemIdResult, error)
	Get(req ItemIdTodoIdRequestUri) (*ItemResult, error)
	Delete(req ItemIdTodoIdRequestUri) (*ItemIdResult, error)
}
