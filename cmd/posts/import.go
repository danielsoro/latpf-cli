package posts

import (
	"log/slog"

	"github.com/danielsoro/wordpress-cli/lib/wordpress"
	"github.com/spf13/cobra"
)

var importCommand = &cobra.Command{
	Use:   "import",
	Short: "Import posts from file to wordpress",
	RunE: func(cmd *cobra.Command, args []string) error {
		var postType wordpress.PostType = "public"
		postFlag, err := cmd.Flags().GetString("type")
		if err != nil {
			slog.Error("post import error", slog.String("error", err.Error()))
			return err
		}

		err = postType.Set(postFlag)
		if err != nil {
			slog.Error("post import error", slog.String("error", err.Error()))
			return err
		}
		return nil
	},
}

func NewImportCommand() *cobra.Command {
	return importCommand
}

func init() {
	importCommand.Flags().StringP("file", "f", "", "The post's file to import.")
	importCommand.Flags().StringP("type", "t", "public", `The post types that should be imported. Allowed: "public", "premium", "curation"`)

	err := importCommand.MarkFlagRequired("file")
	if err != nil {
		slog.Debug("Not possible to mark file flag as required", slog.String("error", err.Error()))
	}
}
