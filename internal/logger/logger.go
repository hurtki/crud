package logger

import (
	"log/slog"
	"os"
)

func NewLogger(level slog.Level) *slog.Logger {
	handlerOptions := &slog.HandlerOptions{
		Level: level,
	}
	handler := slog.NewTextHandler(os.Stdout, handlerOptions)
	logger := slog.New(handler)

	return logger
}
