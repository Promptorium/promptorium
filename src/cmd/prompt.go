/*
Copyright © 2024 Vladislav Parfeniuc
*/
package cmd

import (
	"fmt"
	"promptorium-go/modules/promptmodule"

	"github.com/spf13/cobra"
)

var configPath string
var themePath string
var shell string
var exitCode int

// promptCmd represents the prompt command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {

	fmt.Print(promptmodule.GetPrompt(configPath, themePath, shell, exitCode))
}
