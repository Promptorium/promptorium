#!/bin/bash
function prompt_cmd() {
    # Keep this line at the top of the function, otherwise the exit code will get
    local exit_code="$?"
    local promptorium_output
    promptorium_output=$(promptorium prompt --shell bash --exit-code "$exit_code")
    PS1="$promptorium_output"
}

PROMPT_COMMAND=prompt_cmd