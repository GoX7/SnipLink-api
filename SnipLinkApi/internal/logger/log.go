package logger

import (
	"api/internal/config"
	"log/slog"
	"os"
	"strings"
)

type Logs struct {
	Server *slog.Logger
	Sqlite *slog.Logger
	MW     *slog.Logger
}

func New(cfg *config.Config) (*Logs, error) {
	var level slog.Level
	var logs Logs

	switch cfg.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Create dir "logs"
	dir := strings.Split(cfg.Path.Server, "/")
	os.Mkdir(dir[0], 0755)

	// Open file to write log
	file1, err := os.OpenFile(cfg.Path.Server, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	file2, err := os.OpenFile(cfg.Path.Sqlite, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	file3, err := os.OpenFile(cfg.Path.MW, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	// Register loggers
	logs.Server = slog.New(slog.NewTextHandler(file1, &slog.HandlerOptions{
		Level: level,
	}))
	logs.Sqlite = slog.New(slog.NewTextHandler(file2, &slog.HandlerOptions{
		Level: level,
	}))
	logs.MW = slog.New(slog.NewTextHandler(file3, &slog.HandlerOptions{
		Level: level,
	}))

	return &logs, nil
}
