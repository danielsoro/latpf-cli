package posts

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/danielsoro/wordpress-cli/lib/wordpress/client"
	"github.com/dromara/carbon/v2"
	"github.com/spf13/cobra"
	"strconv"
)

type CreatePost struct {
	Command *cobra.Command
}

func NewCreateCommand(clientType client.WordPressClientType) CreatePost {
	createPostCommand := CreatePost{
		Command: &cobra.Command{
			Use:   "create",
			Short: "Create new wordpress post",
			RunE: func(cmd *cobra.Command, args []string) error {
				wordPressClient, err := client.NewWordPressClient(clientType)
				if err != nil {
					return err
				}

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

				post, err := wordPressClient.CreatePost(title, content, status, carbon.Now().ToRfc3339String())
				if err != nil {
					return err
				}

				rows := [][]string{{strconv.Itoa(post.ID), post.Title.Rendered, post.Content.Raw, post.Link, post.Date}}
				t := table.New().
					Border(lipgloss.NormalBorder()).
					BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
					Headers("ID", "Title", "Content", "Link", "Date")

				t.Rows(rows...)
				fmt.Println(t)
				return nil
			},
		},
	}

	createPostCommand.Command.Flags().StringP("title", "t", "", "The post's title.")
	createPostCommand.Command.Flags().StringP("content", "c", "", "The post's content.")
	createPostCommand.Command.Flags().StringP("status", "s", "publish", "The post's status.")

	return createPostCommand
}
