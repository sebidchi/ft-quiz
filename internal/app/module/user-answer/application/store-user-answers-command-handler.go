package application

import (
	"context"

	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
)

type StoreUserAnswersCommand struct {
	userId  string
	answers map[string]string
}

func NewStoreUserAnswersCommand(userId string, answers map[string]string) *StoreUserAnswersCommand {
	return &StoreUserAnswersCommand{userId: userId, answers: answers}
}

func (s StoreUserAnswersCommand) Id() string {
	return "StoreUserAnswersCommand"
}

type StoreUserAnswersCommandHandler struct {
	userAnswersStorer *domain.UserAnswersStorer
}

func NewStoreUserAnswersCommandHandler(userAnswersStorer *domain.UserAnswersStorer) *StoreUserAnswersCommandHandler {
	return &StoreUserAnswersCommandHandler{userAnswersStorer: userAnswersStorer}
}

func (s StoreUserAnswersCommandHandler) Handle(ctx context.Context, command bus.Command) error {
	storeUserAnswersCommand, ok := command.(*StoreUserAnswersCommand)
	if !ok {
		return bus.NewCommandNotValid(command.Id())
	}

	return s.userAnswersStorer.Store(ctx, storeUserAnswersCommand.userId, storeUserAnswersCommand.answers)
}
