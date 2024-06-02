package pkg

import (
	"log/slog"
	"os"

	"github.com/m-mizutani/clog"
)

func NewLogger() *slog.Logger {
	handler := clog.New(
		clog.WithWriter(os.Stdout),
		clog.WithColor(true),
		clog.WithSource(true),
	)
	logger := slog.New(handler)

	return logger
}
