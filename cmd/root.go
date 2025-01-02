package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
)

type RootFlags struct {
	Username     string
	Password     string
	WordpressURL string
}

var rootCommand = &cobra.Command{
	Use:   "wordpress-cli",
	Short: "wordpress-cli is a tool to admin wordpess via terminal",
}

func init() {
	rootCommand.PersistentFlags().StringP("username", "u", "", "--username my_user")
	rootCommand.PersistentFlags().StringP("password", "p", "", "--password my_password")
	rootCommand.PersistentFlags().StringP("wordpress-url", "", "", "--wordpress-url https://$DOMAIN/wp-json/wp/v2")
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		slog.Error("Not possible to start CLI")
	}
}
