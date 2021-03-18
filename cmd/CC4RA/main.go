package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	apiInternal "github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/internal/api"
	loggerInternal "github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/internal/logger"
	specificationInternal "github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/internal/specification"
	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

func main() {
	var err error

	if err = loggerInternal.Initialize(); err != nil {
		panic(err)
	}
	defer logger.Logger.Sync()

	if err = specificationInternal.Initialize("app"); err != nil {
		panic(err)
	}

	if err = apiInternal.InitializePublic("public_api"); err != nil {
		panic(err)
	}
	logger.Sugar.Debugf("Started public api")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	logger.Sugar.Info("Termination signal received. Starting graceful shutdown.")

	ctx, cancel := context.WithTimeout(context.Background(), specificationInternal.Specification.GracefulShutdownTimeout)
	defer cancel()

	apiInternal.PublicAPI.Server.Shutdown(ctx)
}
