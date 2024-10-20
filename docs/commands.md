---
sidebar_position: 2
---

# Commands

Promptorium provides the following commands:

- `promptorium init`: Initialize Promptorium
- `promptorium shell`: Print the shell script for the current shell
- `promptorium prompt`: Print the promptorium prompt

Ideally, the only command you need to use is `promptorium init` as the other commands are used internally by promptorium. However, you can use the other commands if you want to do something specific.

## Global Flags
- `-d, --debug`: Enable debug mode
- `    --version`: Show version

## promptorium init

This command is used to initialize Promptorium. It generates the following file structure:

```bash
~/.config/promptorium/
├── conf.json
└── presets
    ├── default_1
    │   ├── conf.json
    │   └── theme.json
    └── default_2
        ├── conf.json
        └── theme.json
```
It copies the `conf.json` file from `/usr/share/promptorium/conf` to the `~/.config/promptorium` directory, but only if it doesn't already exist.It also copies the `presets` directory from `/usr/share/promptorium/conf` to the `~/.config/promptorium` directory, but only if the directory doesn't already exist.

Additionally, it adds the line `if [[ $(command -v promptorium) 2> /dev/null ]]; then source <(promptorium shell); fi` to your shell configuration file.

## promptorium shell

This command is used to print the shell script for the current shell.:

Here's the shell script for bash:
```bash
#!/bin/bash
		function prompt_cmd() {
		local exit_code promptorium_output config_file theme_file
		exit_code="$?"
		config_file=$configPath 
		theme_file=$themePath  
		promptorium_output=$(promptorium prompt --shell bash --config-file "$config_file" --theme-file "$theme_file" --exit-code "$exit_code")
		PS1="$promptorium_output"
	}
	PROMPT_COMMAND=prompt_cmd
```

Here's the shell script for zsh:

```zsh

	function set_prompt() {
		local exit_code promptorium_output config_file theme_file
		exit_code="$?"
		config_file=$configPath
		theme_file=$themePath
		promptorium_output=$(promptorium prompt --shell zsh --config-file "$config_file" --theme-file "$theme_file" --exit-code "$exit_code")
		PROMPT="$promptorium_output"
	}
	precmd_functions+=set_prompt
```

Where `$themePath` and `$configPath` are the parameters passed to the `promptorium shell` command.

By default, `promptorium shell` tries to identify the shell using the `SHELL` environment variable. You can specify the shell using the `--shell` flag.

### Flags

- `--config-file`: The path to the config file
- `--theme-file`: The path to the theme file

## promptorium prompt

This command is used to print the promptorium prompt. It formats the prompt based on the configuration file and theme file.


### Flags

- `--shell`: The shell to use (bash, zsh). Default is `bash`
- `--config-file`: The path to the config file
- `--theme-file`: The path to the theme file
- `--exit-code`: The exit code of the last command