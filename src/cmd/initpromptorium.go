package cmd

import (
	"promptorium/internal/pkg/initpkg"

	"github.com/spf13/cobra"
)

var initPromptoriumCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize promptorium",
	Long: `Initialize promptorium by doing the following:
	- Create ~/.config/promptorium directory
	- Copy config files from /usr/share/promptorium/conf to ~/.config/promptorium
	- Copy preset files from /usr/share/promptorium/conf/presets to ~/.config/promptorium/presets
	- Give file permissions to user (might ask for user password)
	- Add line to ~/.bashrc and/or ~/.zshrc to source promptorium shell`,
	Run: func(cmd *cobra.Command, args []string) {
		runInitPromptoriumCmd()
	},
}

func runInitPromptoriumCmd() {
	initpkg.InitPromptorium()
}

func init() {
	rootCmd.AddCommand(initPromptoriumCmd)
}
