package cmd

import (
	"github.com/danielsoro/wordpress-cli/cmd/imports"
	"github.com/spf13/cobra"
	"log/slog"
)

var importCommand = &cobra.Command{
	Use:   "post",
	Short: "Manage post from wordpress",
}

func init() {
	importCommand.AddCommand(imports.NewCurationPostsCommand())
	importCommand.AddCommand(imports.NewCurationPremiumCommand())
	importCommand.AddCommand(imports.NewPublicPostsCommand())

	importCommand.PersistentFlags().StringP("file", "f", "", "--file my_file.cvs")

	err := importCommand.MarkPersistentFlagRequired("file")
	if err != nil {
		slog.Debug("Not possible to mark file flag as required", err)
	}

	rootCommand.AddCommand(importCommand)
}
