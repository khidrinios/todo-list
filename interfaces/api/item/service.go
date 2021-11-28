package item

type Service interface {
	AddItemToTodo(reqParam TodoIdRequestUri, reqbody AddItemToTodoRequestBody) (*AddItemToTodoResult, error)
	GetItemByIdAndTodoId(req ItemIdTodoIdRequestUri) (*ItemResult, error)
}
