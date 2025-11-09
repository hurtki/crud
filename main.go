package main

import (
	"time"

	"github.com/hurtki/crud/internal/db"
	"github.com/hurtki/crud/internal/logger"
)

func main() {
	logger := logger.NewLogger()

	logger.Info("logger initialized")

	db.ConnectDataBase()

	time.Sleep(30 * time.Second)
	
}
