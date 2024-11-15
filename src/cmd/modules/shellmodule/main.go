package shellmodule

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func GetShellScript(shell string, configPath string, themePath string) string {

	if shell == "" {
		shell = filepath.Base(os.Getenv("SHELL"))
	}

	log.Debug().Msgf("[SHELL@shellmodule] Shell: %s", shell)
	return getShellScript(configPath, themePath, shell)
}

func getShellScript(configPath string, themePath string, shell string) string {

	// Check if files exist
	if configPath != "" {
		_, err := os.Stat(configPath)
		if err != nil {
			fmt.Printf("Config file not found: %s", configPath)
			return ""
		}
	}

	if themePath != "" {
		_, err := os.Stat(themePath)
		if err != nil {
			fmt.Printf("Theme file not found: %s", themePath)
			return ""
		}
	}

	switch shell {
	case "bash":
		return getBashScript(configPath, themePath)
	case "zsh":
		return getZshScript(configPath, themePath)
	default:
		return ""
	}
}

func getBashScript(configPath string, themePath string) string {
	bashScript := `
	#!/bin/bash
		function prompt_cmd() {
		local exit_code="$?"
		local promptorium_output config_file theme_file
		config_file=` + configPath + `
		theme_file=` + themePath + `
		promptorium_output=$(promptorium prompt --shell bash --config-file "$config_file" --theme-file "$theme_file" --exit-code "$exit_code")
		PS1="$promptorium_output"
	}
	PROMPT_COMMAND=prompt_cmd`
	return bashScript
}

func getZshScript(configPath string, themePath string) string {
	zshScript := `
	function set_prompt() {
		local exit_code="$?"
		local promptorium_output config_file theme_file
		config_file=` + configPath + `
		theme_file=` + themePath + `
		promptorium_output=$(promptorium prompt --shell zsh --config-file "$config_file" --theme-file "$theme_file" --exit-code "$exit_code")
		PROMPT="$promptorium_output"
	}
	precmd_functions+=set_prompt`
	return zshScript
}
