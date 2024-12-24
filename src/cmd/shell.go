package cmd

import (
	"fmt"
	"promptorium/internal/pkg/shellpkg"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Print the shell script",
	Long:  `Prints the shell script to source in the shell's configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		runShellCmd(cmd.Flags())
	},
}

func init() {
	shellCmd.Flags().StringP("config-file", "c", "", "Path to the config file")
	shellCmd.Flags().StringP("shell", "s", "", "Shell for which to format the prompt (bash, zsh)")
	rootCmd.AddCommand(shellCmd)
}

func runShellCmd(pflags *pflag.FlagSet) {
	var configPath string
	var shell string

	pflags.VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "config-file" {
			configPath = flag.Value.String()
		}
		if flag.Name == "shell" {
			shell = flag.Value.String()
		}
	})

	fmt.Print(shellpkg.GetShellScript(shell, configPath))
}
