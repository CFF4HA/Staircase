package core

import (
	"log"
	"log/slog"
)

var (
	BackendURL  string
	TemplateDir string
	Logger      = slog.New(slog.NewJSONHandler(log.Writer(), &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
)
