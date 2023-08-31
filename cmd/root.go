package cmd

import (
	"fmt"
	"os"

	"battlebit_admin_panel/internal/version"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		logrus.Panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:     "battlebit_admin_panel",
	Short:   "EliteHunterZ battlebit manager",
	Version: version.FullName(),
	Long:    `EliteHunterZ battlebit manager`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute executes the root command and is the entry point of the application from main.go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
