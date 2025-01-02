package cmd

import (
	"github.com/danielsoro/wordpress-cli/cmd/posts"
	"github.com/spf13/cobra"
)

var postCommand = &cobra.Command{
	Use:   "post",
	Short: "Manage post from wordpress",
}

func init() {
	postCommand.AddCommand(posts.NewImportCommand())
	rootCommand.AddCommand(postCommand)
}
