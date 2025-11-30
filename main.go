package main

import (
	"context"
	"os"
	"os/signal"

	tasksHandler "github.com/hurtki/crud/internal/app/tasks"
	"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/crud/internal/server"
	"github.com/hurtki/routego"

	//"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/crud/internal/domain/tasks"
	tasks_repo "github.com/hurtki/crud/internal/repo/tasks"

	"github.com/hurtki/crud/internal/logger"
)

func main() {
	logger := logger.NewLogger()
	// config := config.NewAppConfig(":80", tasksPerPageCount, time.Second*5)
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		logger.Error("can't load config, exiting", "err", err)
	}

	logger.Info("logger initialized")

	storage, err := tasks_repo.GetTaskStorage(*logger)

	if err != nil {
		logger.Error("Can't initialize storage, exiting", "err", err)
		os.Exit(0)
	}

	tasksUseCases := tasks.NewTaskUseCases(&storage, config)

	tasksHandler := tasksHandler.NewTasksHandler(*logger, tasksUseCases)

	routeSet := routego.NewRouteSet()
	routeSet.Add("/tasks/{num}", tasksHandler.ServeReadUpdateDelete)
	routeSet.Add("/tasks/", tasksHandler.ServeCreateList)
	router := routego.NewRouter(routeSet)

	srv := server.NewServer(&router, *config)

	srvErrChan := make(chan error)
	srv.Start(srvErrChan)
	logger.Info("Server started")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer signal.Stop(signalChan)

	select {
	case err = <-srvErrChan:
		logger.Error("error occured in server", "err", err)
	case <-signalChan:
		logger.Info("Stopping server...")
		err = srv.Stop()
		if err != nil {
			if err == context.DeadlineExceeded {
				logger.Error("time left, server didn't success to close all the connections in time")
			} else {
				logger.Error("error while stopping server", "err", err)
			}
		}
		logger.Info("Server stopped")
	}
}
