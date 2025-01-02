package imports

import "github.com/spf13/cobra"

func NewPublicPostsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "public",
		Short: "Import public posts from a file",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
