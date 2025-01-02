package imports

import "github.com/spf13/cobra"

func NewCurationPostsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "curation",
		Short: "Import curation posts from a file",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
