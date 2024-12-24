package shellpkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func GetShellScript(shell string, configPath string) string {

	if shell == "" {
		shell = filepath.Base(os.Getenv("SHELL"))
	}

	log.Trace().Msgf("Shell: %s", shell)
	return getShellScript(configPath, shell)
}

func getShellScript(configPath string, shell string) string {

	// Check if files exist
	if configPath != "" {
		_, err := os.Stat(configPath)
		if err != nil {
			fmt.Printf("Config file not found: %s", configPath)
			return ""
		}
	}

	switch shell {
	case "bash":
		return getBashScript(configPath)
	case "zsh":
		return getZshScript(configPath)
	default:
		return ""
	}
}

func getBashScript(configPath string) string {
	bashScript := `
	#!/bin/bash
		function prompt_cmd() {
		local exit_code="$?"
		local promptorium_output config_file
		config_file=` + configPath + `
		promptorium_output=$(promptorium prompt --shell bash --config-file "$config_file" --exit-code "$exit_code")
		PS1="$promptorium_output"
	}
	PROMPT_COMMAND=prompt_cmd
	source /etc/bash_completion
	source <(promptorium completion bash)`
	return bashScript
}

func getZshScript(configPath string) string {
	zshScript := `
	function set_prompt() {
		local exit_code="$?"
		local promptorium_output config_file
		config_file=` + configPath + `
		promptorium_output=$(promptorium prompt --shell zsh --config-file "$config_file" --exit-code "$exit_code")
		PROMPT="$promptorium_output"
	}
	precmd_functions+=set_prompt
	source <(promptorium completion zsh)`
	return zshScript
}
