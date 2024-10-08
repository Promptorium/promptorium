#!/bin/bash
function prompt_cmd() {
    # Keep this line at the top of the function, otherwise the exit code will get
    local exit_code="$?"
    local promptorium_config_path="$HOME/.config/promptorium/conf.json"
    local promptorium_theme_path="$HOME/.config/promptorium/theme.json"
    local promptorium_output
    promptorium_output=$(promptorium prompt --config "$promptorium_config_path" --theme "$promptorium_theme_path" --shell bash --exit-code "$exit_code")
    PS1="$promptorium_output"
}

PROMPT_COMMAND=prompt_cmd