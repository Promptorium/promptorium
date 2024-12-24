---
sidebar_position: 4
---

# Changelog

## 0.1.0

- Initial release


## 0.1.1

Features:
- Added support for zsh
- Added git status module

## 0.1.2

Features:
- Added os icon module
- Added $exit_code_color variable

Bug Fixes:
- Fixed a bug where the exit code was not being set correctly

## 0.1.3

Improvements:
- Set up CI/CD pipeline
- Added documentation
- Added changelog

## 0.1.4

Improvements:
- Expanded documentation
- Added CI/CD workflows for automatic changelog updating

## 0.1.5

Website:
- Created new website (https://www.promptorium.org) for documentation
- Updated CI/CD pipeline to automatically update the website on every release
- Added new documentation page dedicated to configuration

Installation:
- Improved installation instructions
- Added universal installation script

Commands:
- Added `promptorium shell` command which prints the shell script for the current shell.
- Added `promptorium init` command which does the following:
    - Creates `~/.config/promptorium` directory
    - Copies `conf.json` file from `/usr/share/promptorium/conf` to `~/.config/promptorium` if it doesn't exist
    - Copies `presets` directory from `/usr/share/promptorium/conf` to `~/.config/promptorium` if it doesn't exist
    - Appends `if [[ $(command -v promptorium) 2> /dev/null ]]; then source<(promptorium shell)` to `~/.bashrc` and/or `~/.zshrc` to source promptorium shell script

Fixes:
- Fixed separator being printed when there are no components on the right side
- Added fallback to default config if config file is not found (same for theme file)
- Fixed theme's component dividers being ignored.
- Improved deb package installation, added `scripts/deb/conffiles`

Other changes:
- Changed `--config` flag to `--config-file`
- Changed `--theme` flag to `--theme-file`

## 0.1.6

Breaking changes:
- Renamed `git_color_no_branch` to `git_color_no_repository` in the config file
- Renamed `git_color_no_remote` to `git_color_no_upstream` in the config file

Performance improvements:
- Added lazy loading and caching of all context data (git status, os, shell, etc.)
- Improved git status data retrieval

Overall improvements:
- Added completion for bash and zsh
- Improved debug logging

Refactorings:
- Added `context` package to confmodule, which is responsible for retrieving context data
- Refactored the way modules are loaded and executed, this should make it easier to add new modules and potentially load them on the fly

Fixed bugs:
- Fixed a bug where the exit code was not being correctly set


## 0.1.7

Breaking changes:
- Drastic changes to config file format, see [the website](www.promptorium.org) for more info
- Removed the `--theme-file` flag, it is not needed with new config file format

New features:
- Added support for yaml config files

Config file format changes:
- New config file format:
    - `prompt` - array of strings, representing the prompt. Each string can be a component name or a module name (New in 0.1.7)
        - The prompt can also be defined as an array of arrays, representing a multiline prompt
        - `---` - indicates the **spacer** (separating the left and right sides of the prompt)
            - Each line of a multiline prompt can only contain one **spacer**
    - `theme` - a map of theme settings (already supported in 0.1.6)
    - `components` - a list of components (already supported in 0.1.6)
    - `options` - a map of options (New in 0.1.7)

- All the individual configuration fields can now be loaded from an external file, see [the documentation](www.promptorium.org) for more info
- Components now have a `type` field, which is used to determine the type of the component (module, plugin, text)
- Component dividers are no longer displayed if the background color is set to `transparent` 

Performance improvements:
- Context data is now loaded in parallel, which leads to much faster prompt rendering (now averages at ~ 10 ms)
