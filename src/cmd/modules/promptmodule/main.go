package promptmodule

import (
	"promptorium/cmd/modules/confmodule"

	"github.com/rs/zerolog/log"
	"golang.org/x/term"
)

func GetPrompt(configPath string, themePath string, shell string, exitCode int) string {

	config := confmodule.GetConfig(configPath, themePath, shell, exitCode)

	terminalWidth, _, error := term.GetSize(0)
	if error != nil {
		log.Debug().Msg("[PROMPT@promptmodule] Error getting terminal width")
		terminalWidth = 0
	}

	return NewPromptBuilder(config, terminalWidth).BuildPrompt()

}
