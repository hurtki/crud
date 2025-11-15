package main

import (
	"os"

	"github.com/hurtki/crud/internal/app"
	"github.com/hurtki/crud/internal/app/routeSet"
	tasksHandler "github.com/hurtki/crud/internal/app/tasks"
	"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/crud/internal/db"
	"github.com/hurtki/crud/internal/logger"
)

func main() {
	logger := logger.NewLogger()
	config := config.NewAppConfig(":8000")
	logger.Info("logger initialized")

	storage, err := db.GetStorage(*logger)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(0)
	}

	routerSet := routeSet.NewRouteSet()

	tasksHandler := tasksHandler.NewTasksHandler(*logger, config, &storage)
	routerSet.Add("/", tasksHandler.ServeReadUpdateDelete)
	routerSet.Add("/tasks/", tasksHandler.ServeCreateList)

	router := app.NewRouter(*logger, config, routerSet)
	if err := router.StartRouting(); err != nil {
		logger.Error("failed to start server: " + err.Error())
	}
}
