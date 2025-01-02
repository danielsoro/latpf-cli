package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
)

var rootCommand = &cobra.Command{
	Use:   "wordpress-cli",
	Short: "wordpress-cli is a tool to admin wordpess via terminal",
}

func Execute() {
	rootCommand.AddCommand(NewVersionCommand())
	err := rootCommand.Execute()
	if err != nil {
		slog.Error("Not possible to start CLI")
	}
}
