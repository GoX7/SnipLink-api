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
	var logs Logs

	dir := strings.Split(cfg.Path.Server, "/")
	os.Mkdir(dir[0], 0755)

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

	logs.Server = slog.New(slog.NewTextHandler(file1, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	logs.Sqlite = slog.New(slog.NewTextHandler(file2, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	logs.MW = slog.New(slog.NewTextHandler(file3, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return &logs, nil
}
