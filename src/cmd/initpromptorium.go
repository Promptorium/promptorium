package cmd

import (
	"promptorium/modules/initmodule"

	"github.com/spf13/cobra"
)

var initPromptoriumCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize promptorium",
	Long:  `Initialize promptorium`,
	Run: func(cmd *cobra.Command, args []string) {
		runInitPromptoriumCmd()
	},
}

func runInitPromptoriumCmd() {
	initmodule.InitPromptorium()
}

func init() {
	rootCmd.AddCommand(initPromptoriumCmd)
}
