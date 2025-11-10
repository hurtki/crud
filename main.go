package main

import (
	"fmt"
	"os"
	"time"

	tasks_mod "github.com/hurtki/crud/internal/domain/tasks"

	"github.com/hurtki/crud/internal/db"
	"github.com/hurtki/crud/internal/logger"
)

func main() {
	logger := logger.NewLogger()

	logger.Info("logger initialized")

	storage, err := db.GetStorage(*logger)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(0)
	}
	for {
		tasks, err := storage.GetTasks()
		if err != nil {
			logger.Error("cannot get tasks", "error", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println(tasks)

		task := tasks_mod.Task{Name: "alex", Text: "hello world"}

		err = storage.AddTask(task)

		if err != nil {
			logger.Error("cannot add task", "error", err)
			time.Sleep(5 * time.Second)
			continue
		}
		time.Sleep(5 * time.Second)
	}

}
