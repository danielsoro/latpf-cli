package posts

import (
	"errors"
	"fmt"
	"github.com/danielsoro/wordpress-cli/lib/wordpress"
	"github.com/spf13/cobra"
	"log/slog"
)

var importCommand = &cobra.Command{
	Use:   "import",
	Short: "Import posts from file to wordpress",
	RunE: func(cmd *cobra.Command, args []string) error {
		postType, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		postTypeString := wordpress.PostType(postType)
		err = postTypeString.Set(postType)
		if err != nil {
			return errors.New(fmt.Sprintf("The type flag %v", err.Error()))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		postType, _ := cmd.Flags().GetString("type")
		postTypeString := wordpress.PostType(postType)
		_ = postTypeString.Set(postType)
		println(postTypeString)
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
		slog.Debug("Not possible to mark file flag as required", err)
	}

	err = importCommand.MarkFlagRequired("type")
	if err != nil {
		slog.Debug("Not possible to mark file flag as required", err)
	}

}
