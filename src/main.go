package main

import (
	"os"
	"promptorium/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Version string

var Authors = []string{
	"Vladislav Parfeniuc",
}

func main() {
	// Configure logging
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Run the command
	cmd.Execute(Version)
}
