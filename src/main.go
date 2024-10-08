/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"promptorium-go/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const Version = "0.0.5"

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Run the command
	cmd.Execute()
}
