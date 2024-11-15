package promptmodule

import (
	"promptorium/cmd/modules/confmodule"

	"github.com/rs/zerolog/log"
	"golang.org/x/term"
)

func GetPrompt(configPath string, themePath string, shell string, exitCode int, version string) string {
	config := confmodule.GetConfig(configPath, themePath, shell, exitCode, version)

	terminalWidth, _, error := term.GetSize(0)
	if error != nil {
		log.Trace().Msg("Error getting terminal width")
		terminalWidth = 0
	}

	return NewPromptBuilder(config, terminalWidth).BuildPrompt()

}
