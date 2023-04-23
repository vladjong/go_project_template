package logger

import (
	"log/slog"
	"os"
)

const (
	_Error = "error"
	_Debug = "debug"
	_Warn  = "warn"
	_Info  = "info"
)

func New(level string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: mappLevel(level),
	}))
	slog.SetDefault(logger)
}

func mappLevel(level string) slog.Level {
	switch level {
	case _Error:
		return slog.LevelError
	case _Debug:
		return slog.LevelDebug
	case _Warn:
		return slog.LevelWarn
	default:
		return slog.LevelInfo
	}
}
