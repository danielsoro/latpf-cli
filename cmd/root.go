package cmd

import (
	"github.com/danielsoro/wordpress-cli/lib/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
)

var rootCommand = &cobra.Command{
	Use:   "wordpress-cli",
	Short: "wordpress-cli is a tool to admin wordpess via terminal",
}

func init() {
	config.NewConfig()

	rootCommand.PersistentFlags().StringP("config", "c", "~/.config/wordpress-cli/config", "The configuration file to use")
	rootCommand.PersistentFlags().StringP("username", "u", "", "The wordpress username to use the Wordpress Rest API")
	rootCommand.PersistentFlags().StringP("password", "p", "", "The wordpress password to use the Wordpress Rest API")
	rootCommand.PersistentFlags().StringP("wordpress-url", "", "", "The wordpress url REST Api")

	err := viper.BindPFlag("credentials.username", rootCommand.PersistentFlags().Lookup("username"))
	if err != nil {
		slog.Debug("Not possible to bind username flag with config file: ", err)
	}

	err = viper.BindPFlag("credentials.password", rootCommand.PersistentFlags().Lookup("password"))
	if err != nil {
		slog.Debug("Not possible to bind password flag with config file: ", err)
	}

	err = viper.BindPFlag("url", rootCommand.PersistentFlags().Lookup("wordpress-url"))
	if err != nil {
		slog.Debug("Not possible to bind wordpress-url flag with config file: ", err)
	}
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		slog.Debug("Not possible to start CLI", err)
	}
}
