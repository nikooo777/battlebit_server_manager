package config

import (
	"github.com/kz/discordrus"
	"github.com/pludderio/errs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitializeConfiguration inits the base configuration
func InitializeConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Fatalf("Fatal error config file: %s", errors.FullTrace(errors.Err(err)))
	}

	if viper.GetString("discord_hook") != "" {
		initDiscord()
	}
}

// initDiscord initializes the slack connection and posts info level or greater to the set channel.
func initDiscord() {
	discordHook := viper.GetString("discord_hook")
	logrus.AddHook(
		discordrus.NewHook(discordHook,
			logrus.InfoLevel,
			&discordrus.Opts{
				Username: "BattleBitManager",
			}))
}
