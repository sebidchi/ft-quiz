package bus

import "context"

type Command interface {
	Id() string
}

type CommandBus interface {
	Dispatch(ctx context.Context, command Command) error
	RegisterCommand(command Command, handler CommandHandler) error
}

type CommandHandler interface {
	Handle(ctx context.Context, command Command) error
}

type CommandAlreadyRegistered struct {
	message     string
	commandName string
}

func (i CommandAlreadyRegistered) Error() string {
	return i.message
}

func NewCommandAlreadyRegistered(message string, commandName string) CommandAlreadyRegistered {
	return CommandAlreadyRegistered{message: message, commandName: commandName}
}

type CommandNotRegistered struct {
	message     string
	commandName string
}

func (i CommandNotRegistered) Error() string {
	return i.message
}

func NewCommandNotRegistered(message string, commandName string) CommandNotRegistered {
	return CommandNotRegistered{message: message, commandName: commandName}
}

type CommandNotValid struct {
	message string
}

func NewCommandNotValid(message string) *CommandNotValid {
	return &CommandNotValid{message: message}
}

func (i *CommandNotValid) Error() string {
	return i.message
}
