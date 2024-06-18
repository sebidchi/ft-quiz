package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sebidchi/ft-quiz/cmd/di"
	"go.uber.org/zap"
)

func main() {
	ftQuiz := di.Init()
	ftQuiz.Services.Logger.Warn("starting application")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	defer func() {
		ftQuiz.Services.Logger.Warn("stopping application")
		stop()
	}()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", ftQuiz.Config.ServerPort),
		Handler: ftQuiz.Services.Router.Handler(),
	}

	go func() {
		ftQuiz.Services.Logger.Warn(fmt.Sprintf("starting server at :%s", ftQuiz.Config.ServerPort))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		ftQuiz.Services.Logger.Error("server shutdown error", zap.Error(err))
	}
}
