package cmd

import (
	"promptorium/modules/initconfigmodule"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize promptorium",
	Long:  `Initialize promptorium by creating config and theme files, then sourcing them into the shell profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initConfig() {
	initconfigmodule.InitConfig()
}
