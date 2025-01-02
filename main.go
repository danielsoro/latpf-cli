package main

import (
	"log/slog"
	"os"

	"github.com/danielsoro/wordpress-cli/cmd"
)

const WORDPRESS_CLI_CONFIG = "wordrpess/config"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	cmd.Execute()
}
