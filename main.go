package main

import (
	"os"

	tasksHandler "github.com/hurtki/crud/internal/app/tasks"
	"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/routego"

	//"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/crud/internal/domain/tasks"
	tasks_repo "github.com/hurtki/crud/internal/repo/tasks"

	"github.com/hurtki/crud/internal/logger"
)

const (
	tasksPerPageCount int = 4
)

func main() {
	logger := logger.NewLogger()
	config := config.NewAppConfig(":8000", tasksPerPageCount)
	logger.Info("logger initialized")

	storage, err := tasks_repo.GetTaskStorage(*logger)

	if err != nil {
		logger.Info("Can't initialize storage, exiting")
		os.Exit(0)
	}

	tasksUseCases := tasks.NewTaskUseCases(&storage, config)

	routeSet := routego.NewRouteSet()

	tasksHandler := tasksHandler.NewTasksHandler(*logger, tasksUseCases)
	routeSet.Add("/tasks/{num}", tasksHandler.ServeReadUpdateDelete)
	routeSet.Add("/tasks/", tasksHandler.ServeCreateList)

	routegoConfig := routego.NewRoutegoConfig(config.Port)
	router := routego.NewRouter(routegoConfig, routeSet)
	if err := router.StartRouting(); err != nil {
		logger.Error("failed to start server: " + err.Error())
	}
}
