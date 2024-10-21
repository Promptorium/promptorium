#!/bin/zsh
set_prompt() {
    # Keep this line at the top of the function, otherwise the exit code will get lost
    local exit_code="$?"
    local promptorium_output
    promptorium_output=$(promptorium prompt --shell zsh --exit-code "$exit_code")
    PROMPT="$promptorium_output"
}
precmd_functions+=set_prompt