# promptorium

## About

Promptorium is a modular terminal prompt builder that allows you to create custom prompts for your terminal. It supports multiple modules, such as user, cwd, git branch, git status, time, exit code, and more. You can also customize the appearance of each module using colors and icons.

## Showcase

Promptorium comes with some presets that you can use as a starting point for your prompt. Here are some examples of the default presets:

Preset 1:

![default_1 preset 1](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-1.png)

![default_1 preset 2](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-2.png)

![default_1 preset 4](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-4.png)

![default_1 preset 3](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-3.png)

Preset 2:

![default_2 preset 1](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-5.png)

![default_2 preset 2](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-6.png)




## Prerequisites

- A terminal emulator that supports ANSI escape codes
- A patched font that supports powerline symbols ( [here](https://github.com/powerline/fonts) is a list of patched fonts)

## Installation


### Debian/Ubuntu

You can add the repository and install promptorium using the following commands:

```bash
curl -s http://apt.promptorium.org/gpg-key.public | sudo tee /etc/apt/keyrings/promptorium-gpg.public
echo "deb [arch=amd64 signed-by=/etc/apt/keyrings/promptorium-gpg.public] http://apt.promptorium.org/ unstable main" | sudo tee /etc/apt/sources.list.d/promptorium.list
sudo apt update
sudo apt install promptorium
```

Or you can download the deb package from the [releases page](https://github.com/Promptorium/promptorium/releases) and install it using the following command:

```bash
sudo dpkg -i promptorium_[version]-1_[arch].deb
```

## Usage

Promptorium adds a line to your .bashrc or .zshrc file, depending on your shell. You can run the following command to see the promptorium prompt:

```bash
promptorium prompt
```

You can also pass the following arguments to the prompt command:

- --config: Path to the config file
- --theme: Path to the theme file
- --shell: Shell to use (bash, zsh)
- --exit-code: Exit code of the previous command (default: 0)

For example, to use the default config and theme files and the bash shell, you can run the following command:

```bash
promptorium prompt --config /usr/share/promptorium/conf/conf.json --theme /usr/share/promptorium/conf/theme.json --shell bash --exit-code $?
```

## Configuration

The configuration file is located at ~.config/promptorium/conf.json. It contains an array of components, each of which represents a module in the prompt. Each component has a name, content, and style. The content is the actual content of the module, while the style is used to customize the appearance of the module.

Here is an example of a configuration file:

```json
{
	"components" : [
		{
			"name" : "user_component",
			"content": {
				"module": "user",
				"icon": ""
			},
			"style" : {
				"background_color" : "$primary",
				"foreground_color" : "black"
			}
		},
		{
			"name" : "cwd_component",
			 "content": {
				"module": "cwd",
				"icon": ""
			},
			"style" : {
				"background_color": "$primary",
				"foreground_color": "black",
				"margin" :"1"
			}
		},
		{
			"name" : "git_branch_component",
			 "content": {
				"module": "git_branch",
				"icon": ""
			},
			"style" : {
				"background_color": "transparent",
				"foreground_color": "$git_status_color",
				"start_divider" : "",
				"end_divider" : "",
				"padding": "0",
				"margin":"1 0",
				"align": "right"
			}
		},
		{
			"name" : "git_status_component",
			 "content": {
				"module": "git_status"
			},
			"style" : {
				"background_color": "black",
				"padding": "0",
				"margin": "0 1",
				"align": "left"
			}
		},
		{
			"name" : "time_component",
			"content": {
				"module": "time",
				"icon": ""
				
			},
			"style" : {
				"background_color": "$primary",
				"foreground_color": "black",
				"align": "right"
			}
		},
		{
			"name" :"os_icon_component",
			"content": {
				"module": "os_icon"
			},
			"style" : {
				"background_color": "$exit_code_color",
				"foreground_color": "black",
				"margin": "1",
				"padding": "0",
				"align": "right"
			}
		}
	
	]
}
```

In this example, there are five components:

- user_component: The user's username
- cwd_component: The current working directory
- git_branch_component: The current git branch
- git_status_component: The git status (staged changes, unstaged changes, ahead/behind indicators)
- time_component: The current time

Each component has a style that is used to customize its appearance. The style contains the following properties:

- background_color: The background color of the component
- foreground_color: The foreground color of the component
- start_divider: The start divider of the component
- end_divider: The end divider of the component
- margin: The margin of the component
- padding: The padding of the component
- align: The alignment of the component

The background_color and foreground_color properties can be set to one of the following values:

- "$default": The default color of the component
- "$primary": The primary color of the theme
- "$secondary": The secondary color of the theme
- "$tertiary": The tertiary color of the theme
- "$quaternary": The quaternary color of the theme
- "$success": The success color of the theme
- "$warning": The warning color of the theme
- "$error": The error color of the theme
- "$git_status_color": The git status color of the theme
- "$exit_code_color": The exit code color of the theme

The start_divider and end_divider properties can be set to a string that represents the start and end divider of the component.

The align property can be set to one of the following values:

- "left": Align the component to the left
- "right": Align the component to the right

The margin and padding properties can be set to a number or a string that represents the margin or padding in pixels.

## Themes

The theme file is located at ~/.config/promptorium/theme.json.

- component_start_divider: The start divider of the component
- component_end_divider: The end divider of the component
- spacer: The spacer between components
- primary_color: The primary color of the theme
- secondary_color: The secondary color of the theme
- tertiary_color: The tertiary color of the theme
- quaternary_color: The quaternary color of the theme
- success_color: The success color of the theme
- warning_color: The warning color of the theme
- error_color: The error color of the theme
- background_color: The background color of the theme
- foreground_color: The foreground color of the theme
- git_status_clean: The color of the clean git status
- git_status_dirty: The color of the dirty git status
- git_status_no_branch: The color of the git status when there is no branch
- git_status_no_remote: The color of the git status when there is no remote
- exit_code_color_ok: The color of the exit code when the exit code is 0
- exit_code_color_error: The color of the exit code when the exit code is not 0

## Presets


You can also specify a preset in the configuration file. If you do, promptorium will look for a configuration file in /usr/share/promptorium/presets/[preset name]/conf.json and a theme file in /usr/share/promptorium/presets/[preset name]/theme.json.
When you specify a preset, all other properties in the configuration file will be ignored.

Promptorium comes with two presets: default_1 and default_2, but you can create your own presets by creating a new directory in /usr/share/promptorium/presets/ and adding a conf.json and a theme.json file to it.