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