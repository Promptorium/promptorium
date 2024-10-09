package cmd

import (
	"fmt"
	"promptorium/modules/promptmodule"

	"github.com/spf13/cobra"
)

var configPath string
var themePath string
var shell string
var exitCode int

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Print the prompt",
	Long: `Prints the prompt string based on the config file.
	If the config file is not found, it will print a default prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	promptCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to the config file")
	promptCmd.Flags().StringVarP(&themePath, "theme", "t", "", "Path to the theme file")
	promptCmd.Flags().StringVarP(&shell, "shell", "s", "", "Shell to use (bash, zsh)")
	promptCmd.Flags().IntVarP(&exitCode, "exit-code", "e", 0, "Exit code of the previous command")
	rootCmd.AddCommand(promptCmd)
}

func run() {

	fmt.Print(promptmodule.GetPrompt(configPath, themePath, shell, exitCode))
}
