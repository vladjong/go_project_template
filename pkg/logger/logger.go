package logger

import (
	"log/slog"
	"os"
)

const (
	_error = "error"
	_debug = "debug"
	_warn  = "warn"
	_info  = "info"
)

func New(level string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: mappLevel(level),
	}))
	slog.SetDefault(logger)
}

func mappLevel(level string) slog.Level {
	switch level {
	case _error:
		return slog.LevelError
	case _debug:
		return slog.LevelDebug
	case _warn:
		return slog.LevelWarn
	default:
		return slog.LevelInfo
	}
}
