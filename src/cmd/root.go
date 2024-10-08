/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "promptorium",
	Short: "A modular terminal prompt builder",
	Long:  `A modular terminal prompt builder`,
}

func Execute() {
	// Set log level before running the command
	rootCmd.ParseFlags(os.Args)
	if rootCmd.Flags().Changed("debug") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug mode enabled")
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug mode")
}
