package cmd

import (
	"battlebit_admin_panel/internal/version"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of battlebit manager",
	Long:  `All software have versions. This is ours`,
	Run: func(cmd *cobra.Command, args []string) {
		println(version.FullName())
	},
}
