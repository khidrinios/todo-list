package todo

type Service interface {
	Create(req CreateRequest) (*TodoIdResult, error)
	GetById(req TodoByIdRequestUri) (*TodoResult, error)
	List(req QueryTodosRequestBody) ([]TodoResult, error)
	DeleteById(req TodoByIdRequestUri) (*TodoIdResult, error)
	Update(reqParam TodoByIdRequestUri, reqbody UpdateTodoRequestBody) (*UpdateTodoResult, error)
}
