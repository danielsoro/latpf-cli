package posts

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/danielsoro/wordpress-cli/lib/wordpress"
	"github.com/spf13/cobra"
	"strconv"
)

type post struct {
	Id            int    `json:"id"`
	TitleRaw      string `json:"titleRaw"`
	TitleRendered string `json:"titleRendered"`
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List posts from wordpress",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := wordpress.NewWordpressWithViper()

		posts := client.GetPosts()
		if len(posts) == 0 {
			fmt.Println("No posts found")
			return nil
		}

		rows := [][]string{}

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
			Headers("ID", "Title", "Content", "Link")

		for _, p := range posts {
			rows = append(rows, []string{
				strconv.Itoa(p.ID), p.Title.Rendered, p.Content.Raw, p.Link,
			})
		}

		t.Rows(rows...)

		fmt.Println(t)
		return nil
	},
}

func NewListCommand() *cobra.Command {
	return listCommand
}
