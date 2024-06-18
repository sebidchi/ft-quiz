package di

import (
	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/application"
	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/domain"
	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/infrastructure"
	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/ui"
	config "github.com/sebidchi/ft-quiz/internal/pkg/infrastructure"
)

type quizServices struct {
	FetchQuizQueryHandler *application.FetchQuizQueryHandler
	QuizRepository        domain.QuizRepository
}

func initQuizServices(commonServices *CommonServices, config config.Config) *quizServices {
	repository := infrastructure.NewInMemoryQuizRepository(randomQuizWithQuestions())
	quizFetcher := domain.NewQuizFetcher(repository)
	qs := &quizServices{
		FetchQuizQueryHandler: application.NewFetchQuizQueryHandler(quizFetcher),
		QuizRepository:        repository,
	}

	registerQuizRoutes(commonServices, qs)
	registerQuizQueryHandlers(commonServices, qs)

	return qs
}

func registerQuizRoutes(commonServices *CommonServices, _ *quizServices) {
	commonServices.Router.GET("/quiz", ui.HandleGetQuiz(commonServices.QueryBus, commonServices.HttpServices.JsonResponseMiddleware))
}

func registerQuizQueryHandlers(commonServices *CommonServices, qs *quizServices) {
	registerQueryHandlerOrPanic(commonServices.QueryBus, &application.FetchQuizQuery{}, qs.FetchQuizQueryHandler)
}

func randomQuizWithQuestions() *domain.Quiz {
	questions := []domain.Question{
		domain.NewQuestion(
			"question1",
			"What is the capital of France?",
			[]domain.Option{
				*domain.NewOption("option1", "Paris"),
				*domain.NewOption("option2", "Berlin"),
				*domain.NewOption("option3", "London"),
				*domain.NewOption("option4", "Madrid"),
			},
			"option1",
		),
		domain.NewQuestion(
			"question2",
			"Who wrote the novel '1984'?",
			[]domain.Option{
				*domain.NewOption("option5", "George Orwell"),
				*domain.NewOption("option6", "Aldous Huxley"),
				*domain.NewOption("option7", "Ernest Hemingway"),
				*domain.NewOption("option8", "Mark Twain"),
			},
			"option5",
		),
		domain.NewQuestion(
			"question3",
			"What is the square root of 81?",
			[]domain.Option{
				*domain.NewOption("option9", "7"),
				*domain.NewOption("option10", "8"),
				*domain.NewOption("option11", "9"),
				*domain.NewOption("option12", "10"),
			},
			"option11",
		),
		domain.NewQuestion(
			"question4",
			"What is the chemical symbol for Hydrogen?",
			[]domain.Option{
				*domain.NewOption("option13", "H"),
				*domain.NewOption("option14", "He"),
				*domain.NewOption("option15", "Hy"),
				*domain.NewOption("option16", "Ho"),
			},
			"option13",
		),
		domain.NewQuestion(
			"question5",
			"Who painted the Mona Lisa?",
			[]domain.Option{
				*domain.NewOption("option17", "Vincent van Gogh"),
				*domain.NewOption("option18", "Pablo Picasso"),
				*domain.NewOption("option19", "Leonardo da Vinci"),
				*domain.NewOption("option20", "Claude Monet"),
			},
			"option19",
		),
	}

	return domain.NewQuiz("real-quiz-id", "Real Quiz", "A quiz with real questions", questions)
}
