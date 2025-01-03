package main

import (
	"log/slog"
	"os"

	"github.com/danielsoro/wordpress-cli/cmd"
)

func getLogLevel(s string) slog.Level {
	switch s {
	case "debug":
		return slog.LevelDebug
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo

	}
}

func main() {
	logLevel := getLogLevel(os.Getenv("WCLI_LOG"))

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	slog.SetDefault(logger)
	cmd.Execute()
}
