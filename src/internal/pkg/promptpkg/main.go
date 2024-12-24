package promptpkg

import (
	"promptorium/internal/pkg/confpkg/config"
)

func GetPrompt(configPath string, shell string, exitCode int, version string) string {
	config := config.GetConfig(configPath, shell, exitCode, version)
	return NewPromptBuilder(config).BuildPrompt().Render()

}
