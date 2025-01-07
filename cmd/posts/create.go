package posts

import (
	"encoding/json"
	"fmt"
	"github.com/danielsoro/wordpress-cli/lib/wordpress"
	"github.com/spf13/cobra"
)

type CreatePost struct {
	client  wordpress.Wordpress
	Command *cobra.Command
}

func NewImportCommand(client wordpress.Wordpress) CreatePost {
	createPostCommand := CreatePost{
		client: client,
		Command: &cobra.Command{
			Use:   "create",
			Short: "Create new wordpress post",
			RunE: func(cmd *cobra.Command, args []string) error {
				title, err := cmd.Flags().GetString("title")
				if err != nil {
					return err
				}

				content, err := cmd.Flags().GetString("content")
				if err != nil {
					return err
				}

				status, err := cmd.Flags().GetString("status")
				if err != nil {
					return err
				}

				post, err := client.CreatePost(title, content, status)
				if err != nil {
					return err
				}

				result, err := json.MarshalIndent(post, "", " ")
				if err != nil {
					return err
				}

				print(fmt.Sprintf("%s", result))
				return nil
			},
		},
	}

	createPostCommand.Command.Flags().StringP("title", "t", "", "The post's title.")
	createPostCommand.Command.Flags().StringP("content", "c", "", "The post's content.")
	createPostCommand.Command.Flags().StringP("status", "s", "publish", "The post's status.")

	return createPostCommand
}
