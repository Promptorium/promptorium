package shellmodule

import (
	"os"
	"path/filepath"
)

func GetShellScript(shell string, configPath string, themePath string) string {

	if shell == "" {
		shell = filepath.Base(os.Getenv("SHELL"))
	}

	return getShellScript(configPath, themePath, shell)
}

func getShellScript(configPath string, themePath string, shell string) string {

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
		local exit_code promptorium_output config_path theme_path
		exit_code="$?"
		config_path=` + configPath + `
		theme_path=` + themePath + `
		promptorium_output=$(promptorium prompt --shell bash --config-file "$config_path" --theme-file "$theme_path" --exit-code "$exit_code")
		PS1="$promptorium_output"
	}
	PROMPT_COMMAND=prompt_cmd`
	return bashScript
}

func getZshScript(configPath string, themePath string) string {
	zshScript := `
	function set_prompt() {
		local exit_code promptorium_output config_path theme_path
		exit_code="$?"
		config_path=` + configPath + `
		theme_path=` + themePath + `
		promptorium_output=$(promptorium prompt --shell zsh --config-file "$config_path" --theme-file "$theme_path" --exit-code "$exit_code")
		PROMPT="$promptorium_output"
	}
	precmd_functions+=set_prompt`
	return zshScript
}
