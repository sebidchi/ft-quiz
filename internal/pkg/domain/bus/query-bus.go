package bus

import "context"

type Query interface {
	Id() string
}

type QueryBus interface {
	Ask(ctx context.Context, query Query) (interface{}, error)
	RegisterQuery(query Query, handler QueryHandler) error
}

type QueryHandler interface {
	Handle(ctx context.Context, query Query) (interface{}, error)
}

type InvalidQuery struct {
}

func NewInvalidQuery(query Query) *InvalidQuery {
	return &InvalidQuery{}
}

func (i InvalidQuery) Error() string {
	return "Invalid query"
}

type QueryAlreadyRegistered struct {
	message   string
	queryName string
}

func (i QueryAlreadyRegistered) Error() string {
	return i.message
}

func NewQueryAlreadyRegistered(message string, queryName string) QueryAlreadyRegistered {
	return QueryAlreadyRegistered{message: message, queryName: queryName}
}

type QueryNotRegistered struct {
	message   string
	queryName string
}

func (i QueryNotRegistered) Error() string {
	return i.message
}

func NewQueryNotRegistered(message string, queryName string) QueryAlreadyRegistered {
	return QueryAlreadyRegistered{message: message, queryName: queryName}
}
