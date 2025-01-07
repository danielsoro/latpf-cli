package cmd

import (
	"github.com/danielsoro/wordpress-cli/lib/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version of wordpress-cli",
		Run: func(cmd *cobra.Command, args []string) {
			version := version.NewVersion()
			cmd.Println(version.Execute())
		},
	}
}
