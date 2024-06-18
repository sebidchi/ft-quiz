package bus

import (
	"context"
	"reflect"
	"sync"

	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
)

type CommandBus struct {
	handlers map[string]bus.CommandHandler
	lock     sync.Mutex
	wg       sync.WaitGroup
}

func InitCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]bus.CommandHandler, 0),
		lock:     sync.Mutex{},
		wg:       sync.WaitGroup{},
	}
}

func (cb *CommandBus) RegisterCommand(command bus.Command, handler bus.CommandHandler) error {
	cb.lock.Lock()
	defer cb.lock.Unlock()

	commandName, err := cb.commandName(command)
	if err != nil {
		return err
	}

	if _, ok := cb.handlers[*commandName]; ok {
		return bus.NewCommandAlreadyRegistered("Command already registered", *commandName)
	}

	cb.handlers[*commandName] = handler

	return nil
}

func (cb *CommandBus) Dispatch(ctx context.Context, command bus.Command) error {
	commandName, err := cb.commandName(command)
	if err != nil {
		return err
	}

	if handler, ok := cb.handlers[*commandName]; ok {
		err = handler.Handle(ctx, command)
		if err != nil {
			return err
		}

		return nil
	}

	return bus.NewCommandNotRegistered("Command not registered", *commandName)
}

func (cb *CommandBus) commandName(cmd bus.Command) (*string, error) {
	value := reflect.ValueOf(cmd)

	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		return nil, bus.NewCommandNotValid("only pointer to commands are allowed")
	}

	name := value.String()

	return &name, nil
}
