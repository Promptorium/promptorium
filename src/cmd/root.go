package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Version string
var rootCmd = &cobra.Command{
	Use:   "promptorium",
	Short: "A modular terminal prompt builder",
	Long:  `A modular terminal prompt builder`,
}

func Execute() {
	rootCmd.Version = Version
	// Set log level before running the command
	start := time.Now()
	rootCmd.ParseFlags(os.Args)
	if rootCmd.Flags().Changed("debug") {
		switch rootCmd.Flags().Lookup("debug").Value.String() {
		case "1":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			log.Debug().Msg("Debug mode enabled")
		case "2":
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
			log.Trace().Msg("Trace mode enabled")
		default:
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	elapsed := time.Since(start)
	if zerolog.GlobalLevel() == zerolog.DebugLevel || zerolog.GlobalLevel() == zerolog.TraceLevel {
		// add a new line after the prompt if debug mode is enabled
		fmt.Println()
	}
	log.Debug().Msgf("Execution time: %s", elapsed)
}

func init() {
	rootCmd.PersistentFlags().CountP("debug", "d", "Debug mode")
}
