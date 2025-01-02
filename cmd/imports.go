package cmd

import (
	"github.com/danielsoro/wordpress-cli/cmd/imports"
	"github.com/spf13/cobra"
)

var importCommand = &cobra.Command{
	Use:   "imports",
	Short: "Import posts from a file",
}

func NewImportsCommand() *cobra.Command {
	importCommand.AddCommand(imports.NewCurationPostsCommand())
	importCommand.AddCommand(imports.NewCurationPremiumCommand())
	importCommand.AddCommand(imports.NewPublicPostsCommand())

	importCommand.PersistentFlags().StringP("file", "f", "", "--file my_file.cvs")

	return importCommand
}

func init() {
	rootCommand.AddCommand(NewImportsCommand())
}
