package core

import (
	"log/slog"
	"os"
)

var (
	opts = &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
)
