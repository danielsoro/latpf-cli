package main

import (
	"log/slog"
	"os"

	"github.com/danielsoro/wordpress-cli/cmd"
)

func main() {
	logLevel := os.Getenv("WCLI_DEBUG")

	if len(logLevel) > 0 {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		slog.SetDefault(logger)
		cmd.Execute()
		return
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cmd.Execute()
}
