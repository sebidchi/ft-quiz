package di

import (
	"fmt"

	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/application"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/infrastructure"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/ui"
	config "github.com/sebidchi/ft-quiz/internal/pkg/infrastructure"
)

type userAnswersServices struct {
	StoreUserAnswersQueryHandler *application.StoreUserAnswersCommandHandler
	GetUserResultsQueryHandler   *application.GetUserResultsQueryHandler
	UserAnswerRepository         domain.UserAnswerRepository
}

func initUserAnswersServices(commonServices *CommonServices, cnf config.Config) *userAnswersServices {
	repository := infrastructure.NewInMemoryUserAnswerRepository()
	storer := domain.NewUserAnswersStorer(repository)
	uas := &userAnswersServices{
		StoreUserAnswersQueryHandler: application.NewStoreUserAnswersCommandHandler(storer),
		GetUserResultsQueryHandler:   application.NewGetUserResultsQueryHandler(domain.NewUserResultsFetcher(repository)),
		UserAnswerRepository:         repository,
	}

	registerUserAnswersRoutes(commonServices, cnf, uas)
	registerUserAnswersCommandHandlers(commonServices, uas)
	registerUserAnswersQueryHandlers(commonServices, uas)

	return uas
}

func registerUserAnswersRoutes(commonServices *CommonServices, cnf config.Config, _ *userAnswersServices) {
	commonServices.Router.POST(
		"/answers",
		commonServices.HttpServices.JsonSchemaValidator.Validate(fmt.Sprintf("%s%s", cnf.SchemaPath, ui.PostUserAnswersSchema)),
		ui.HandlePostUserAnswers(commonServices.CommandBus, commonServices.HttpServices.JsonResponseMiddleware),
	)
	commonServices.Router.GET("/answers/:username",
		ui.HandleGetUserResults(commonServices.QueryBus, commonServices.HttpServices.JsonResponseMiddleware),
	)
}

func registerUserAnswersCommandHandlers(commonServices *CommonServices, uas *userAnswersServices) {
	registerCommandHandlerOrPanic(commonServices.CommandBus, &application.StoreUserAnswersCommand{}, uas.StoreUserAnswersQueryHandler)
}

func registerUserAnswersQueryHandlers(commonServices *CommonServices, uas *userAnswersServices) {
	registerQueryHandlerOrPanic(commonServices.QueryBus, &application.GetUserResultsQuery{}, uas.GetUserResultsQueryHandler)
}
