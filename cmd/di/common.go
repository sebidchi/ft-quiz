package di

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
	config "github.com/sebidchi/ft-quiz/internal/pkg/infrastructure"
	infraBus "github.com/sebidchi/ft-quiz/internal/pkg/infrastructure/bus"
	"github.com/sebidchi/ft-quiz/pkg/logger"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
)

type CommonServices struct {
	Logger logger.Logger

	CommandBus bus.CommandBus
	QueryBus   bus.QueryBus

	HttpServices *httpServices
	Router       *gin.Engine
}

func Init() *FTQuizDI {
	return setUp()
}

func InitWithEnvFile(envFiles ...string) *FTQuizDI {
	err := godotenv.Overload(envFiles...)
	if err != nil {
		panic(err)
	}

	return setUp()
}

type FTQuizDI struct {
	Services *CommonServices
	Config   config.Config

	QuizServices        *quizServices
	UserAnswersServices *userAnswersServices
}

func setUp() *FTQuizDI {
	cnf := buildConfig()
	l := buildLogger()

	commandBus := infraBus.InitCommandBus()
	queryBus := infraBus.InitQueryBus()

	rt := initRouter(cnf)
	services := &CommonServices{
		Logger:       l,
		CommandBus:   commandBus,
		QueryBus:     queryBus,
		HttpServices: initHttpServices(l),
		Router:       rt,
	}
	ftQuizDI := &FTQuizDI{
		Services:            services,
		Config:              cnf,
		QuizServices:        initQuizServices(services, cnf),
		UserAnswersServices: initUserAnswersServices(services, cnf),
	}

	rt.Use(
		services.HttpServices.RequestPanicMiddleware.RequestPanicHandler(),
		services.HttpServices.RequestIdMiddleware.RequestIdentifier(),
	)

	return ftQuizDI
}

func buildConfig() config.Config {
	var c config.Config
	ctx := context.Background()
	if err := envconfig.Process(ctx, &c); err != nil {
		panic(err)
	}

	return c
}

func buildLogger() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return l
}

func registerQueryHandlerOrPanic(bus bus.QueryBus, query bus.Query, handler bus.QueryHandler) {
	if err := bus.RegisterQuery(query, handler); err != nil {
		panic(err)
	}
}
func registerCommandHandlerOrPanic(bus bus.CommandBus, command bus.Command, handler bus.CommandHandler) {
	if err := bus.RegisterCommand(command, handler); err != nil {
		panic(err)
	}
}

func initRouter(config config.Config) *gin.Engine {
	if config.AppEnv != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	return gin.Default()
}
