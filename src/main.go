package main

import (
	"promptorium/cmd"
	"promptorium/internal/utils"

	"github.com/rs/zerolog/log"
)

var Version string

func main() {
	// Configure logging
	log.Logger = utils.GetLogger()

	cmd.Version = Version
	// Run the command
	cmd.Execute()
}
