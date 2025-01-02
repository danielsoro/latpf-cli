package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "wordpress-cli",
	Short: "wordpress-cli is a tool to admin wordpess via terminal",
}

func Execute() {
	rootCommand.AddCommand(NewVersionCommand())
	rootCommand.Execute()
}
