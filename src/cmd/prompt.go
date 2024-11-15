package cmd

import (
	"fmt"
	"promptorium/cmd/modules/promptmodule"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Print the prompt",
	Long: `Prints the prompt string based on the config file.
	If the config file is not found, it will print a default prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		runPromptCmd(cmd.Flags(), Version)
	},
}

func init() {
	promptCmd.Flags().StringP("config-file", "c", "", "Path to the config file")
	promptCmd.Flags().StringP("theme-file", "t", "", "Path to the theme file")
	promptCmd.Flags().StringP("shell", "s", "", "Shell for which to format the prompt (bash, zsh)")
	promptCmd.Flags().IntP("exit-code", "e", 0, "Exit code of the previous command")
	rootCmd.AddCommand(promptCmd)
}

func runPromptCmd(pFlags *pflag.FlagSet, version string) {

	var configPath string
	var themePath string
	var shell string
	var exitCode int

	pFlags.VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "config-file" {
			configPath = flag.Value.String()
		}
		if flag.Name == "theme-file" {
			themePath = flag.Value.String()
		}
		if flag.Name == "shell" {
			shell = flag.Value.String()
		}
		if flag.Name == "exit-code" {
			exitCode, _ = strconv.Atoi(flag.Value.String())
		}
	})
	log.Debug().Msgf("Version: %s", version)

	fmt.Print(promptmodule.GetPrompt(configPath, themePath, shell, exitCode, version))
}
