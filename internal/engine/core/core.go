package core

import (
	"log"
	"log/slog"
)

type Configuration struct {
	DatabasePollInterval int `json:"database_poll_interval"`
}

var (
	Logger = slog.New(slog.NewJSONHandler(log.Writer(), &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	Config = &Configuration{}
)
