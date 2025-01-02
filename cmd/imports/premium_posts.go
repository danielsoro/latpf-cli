package imports

import "github.com/spf13/cobra"

func NewCurationPremiumCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "premium",
		Short: "Import premium posts from a file",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
