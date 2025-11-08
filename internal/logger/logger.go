package logger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	handlerOptions := &slog.HandlerOptions{
		Level: slog.Level(0),
	}
	handler := slog.NewTextHandler(os.Stdout, handlerOptions)
	logger := slog.New(handler)

	return logger
}
