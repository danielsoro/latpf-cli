package cmd

import (
	"log/slog"

	"github.com/danielsoro/wordpress-cli/lib/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCommand = &cobra.Command{
	Use:   "wordpress-cli",
	Short: "wordpress-cli is a tool to admin wordpess via terminal",
}

func init() {
	config.NewConfig()
	definePersistentFlags()
	bindValuesWithViper()
	makeFlagsRequired()
}

func definePersistentFlags() {
	rootCommand.PersistentFlags().StringP("username", "u", "", "The wordpress username to use the Wordpress Rest API")
	rootCommand.PersistentFlags().StringP("password", "p", "", "The wordpress password to use the Wordpress Rest API")
	rootCommand.PersistentFlags().StringP("wordpress-url", "", "", "The wordpress url REST Api")
}

func bindValuesWithViper() {
	viper.BindPFlag("credentials.username", rootCommand.PersistentFlags().Lookup("username"))
	viper.BindPFlag("credentials.password", rootCommand.PersistentFlags().Lookup("password"))
	viper.BindPFlag("url", rootCommand.PersistentFlags().Lookup("wordpress-url"))
}

func makeFlagsRequired() {
	if len(viper.GetString("url")) == 0 {
		rootCommand.MarkPersistentFlagRequired("wordpress-url")
	}

	if len(viper.GetString("credentials.username")) == 0 {
		rootCommand.MarkPersistentFlagRequired("username")
	}

	if len(viper.GetString("credentials.password")) == 0 {
		rootCommand.MarkPersistentFlagRequired("password")
	}
}

func RemoveRootRequiredFlags(c *cobra.Command) {
	c.InheritedFlags().SetAnnotation("username", cobra.BashCompOneRequiredFlag, []string{"false"})
	c.InheritedFlags().SetAnnotation("password", cobra.BashCompOneRequiredFlag, []string{"false"})
	c.InheritedFlags().SetAnnotation("wordpress-url", cobra.BashCompOneRequiredFlag, []string{"false"})
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		slog.Error("Not possible to start CLI", slog.String("error", err.Error()))
	}
}
