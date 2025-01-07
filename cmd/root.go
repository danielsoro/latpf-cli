package cmd

import (
	"github.com/danielsoro/wordpress-cli/lib/config"
	"github.com/danielsoro/wordpress-cli/lib/wordpress/client"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	config.NewConfig()

	rootCommand := &cobra.Command{
		Use:   "wordpress-cli",
		Short: "wordpress-cli is a tool to admin wordpess via terminal",
	}

	rootCommand.PersistentFlags().StringP("username", "u", "", "The wordpress username to use the Wordpress Rest API")
	rootCommand.PersistentFlags().StringP("password", "p", "", "The wordpress password to use the Wordpress Rest API")
	rootCommand.PersistentFlags().StringP("wordpress-url", "", "", "The wordpress url REST Api")

	_ = viper.BindPFlag("credentials.username", rootCommand.PersistentFlags().Lookup("username"))
	_ = viper.BindPFlag("credentials.password", rootCommand.PersistentFlags().Lookup("password"))
	_ = viper.BindPFlag("url", rootCommand.PersistentFlags().Lookup("wordpress-url"))

	rootCommand.AddCommand(NewVersionCommand())
	rootCommand.AddCommand(NewPostCommand(client.VIPER).Command)
	err := rootCommand.Execute()

	if err != nil {
		slog.Error("Not possible to start CLI", slog.String("error", err.Error()))
	}
}
