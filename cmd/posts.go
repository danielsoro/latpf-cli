package cmd

import (
	"github.com/danielsoro/wordpress-cli/cmd/posts"
	"github.com/danielsoro/wordpress-cli/lib/wordpress"
	"github.com/spf13/cobra"
)

type PostCommand struct {
	client  wordpress.Wordpress
	Command *cobra.Command
}

var postCommand = &cobra.Command{
	Use:   "posts",
	Short: "Manage post from wordpress",
}

func NewPostCommand(clientType wordpress.ClientType) PostCommand {
	postCommand.AddCommand(posts.NewCreateCommand(clientType).Command)
	postCommand.AddCommand(posts.NewListCommand(clientType).Command)

	return PostCommand{
		Command: postCommand,
	}
}
