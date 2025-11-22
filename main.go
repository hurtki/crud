package main

import (
	"os"

	tasksHandler "github.com/hurtki/crud/internal/app/tasks"
	//"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/crud/internal/domain/tasks"
	"github.com/hurtki/crud/internal/repo/tasks"

	"github.com/hurtki/crud/internal/logger"

	"github.com/hurtki/routego"
	"github.com/hurtki/routego/routeSet"

)

func main() {
	logger := logger.NewLogger()
	// TODO: use app config 
	// config := config.NewAppConfig(":8000")
	logger.Info("logger initialized")

	storage, err := tasks_repo.GetTaskStorage(*logger)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(0)
	}

	tasksUseCases := tasks.NewTaskUseCases(&storage)

	routeSet := routeSet.NewRouteSet()

	tasksHandler := tasksHandler.NewTasksHandler(*logger, tasksUseCases)
	routeSet.Add("/tasks/{num}", tasksHandler.ServeReadUpdateDelete)
	routeSet.Add("/tasks/", tasksHandler.ServeCreateList)
	
	routegoConfig := routego.NewRoutegoConfig(":8000")
	router := routego.NewRouter(*logger, routegoConfig, routeSet)
	if err := router.StartRouting(); err != nil {
		logger.Error("failed to start server: " + err.Error())
	}
}
