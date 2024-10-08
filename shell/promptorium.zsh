#!/bin/zsh
set_prompt() {
    # Keep this line at the top of the function, otherwise the exit code will get lost
    local exit_code="$?"
    local promptorium_config_path="$HOME/.config/promptorium/conf.json"
    local promptorium_theme_path="$HOME/.config/promptorium/theme.json"
    local promptorium_output
    promptorium_output=$(promptorium prompt --config "$promptorium_config_path" --theme "$promptorium_theme_path" --shell zsh --exit-code "$exit_code")
    PROMPT="$promptorium_output"
}
precmd_functions+=set_prompt