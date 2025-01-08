package cmd

import (
	"github.com/danielsoro/wordpress-cli/cmd/posts"
	"github.com/danielsoro/wordpress-cli/lib/wordpress/client"
	"github.com/spf13/cobra"
)

type PostCommand struct {
	client  client.Wordpress
	Command *cobra.Command
}

var postCommand = &cobra.Command{
	Use:   "posts",
	Short: "Manage post from wordpress",
}

func NewPostCommand(clientType client.WordPressClientType) PostCommand {
	postCommand.AddCommand(posts.NewCreateCommand(clientType).Command)
	postCommand.AddCommand(posts.NewListCommand(clientType).Command)
	postCommand.AddCommand(posts.NewImportCommand(clientType).Command)

	return PostCommand{
		Command: postCommand,
	}
}
