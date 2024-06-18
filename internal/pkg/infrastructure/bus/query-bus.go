package bus

import (
	"context"
	"sync"

	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
)

type QueryBus struct {
	handlers map[string]bus.QueryHandler
	lock     sync.Mutex
}

func InitQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]bus.QueryHandler, 0),
		lock:     sync.Mutex{},
	}
}

func (qb *QueryBus) RegisterQuery(query bus.Query, handler bus.QueryHandler) error {
	qb.lock.Lock()
	defer qb.lock.Unlock()

	queryName := query.Id()

	if _, ok := qb.handlers[queryName]; ok {
		return bus.NewQueryAlreadyRegistered("Query already registered", queryName)
	}

	qb.handlers[queryName] = handler

	return nil
}

func (qb *QueryBus) Ask(ctx context.Context, query bus.Query) (interface{}, error) {
	queryName := query.Id()

	if handler, ok := qb.handlers[queryName]; ok {
		response, err := qb.doAsk(ctx, handler, query)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	return nil, bus.NewQueryNotRegistered("Query not registered", queryName)
}

func (qb *QueryBus) doAsk(ctx context.Context, handler bus.QueryHandler, query bus.Query) (interface{}, error) {
	return handler.Handle(ctx, query)
}
