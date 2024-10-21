# Promptorium

## About

Promptorium is a modular terminal prompt builder that allows you to create custom prompts for your terminal. It supports multiple modules, such as user, cwd, git branch, git status, time, exit code, and more. You can also customize the appearance of each module using colors and icons.

For more information visit the [documentation website](https://www.promptorium.org)

## Showcase

Promptorium comes with some presets that you can use as a starting point for your prompt. Here are some examples of the default presets:

Preset 1:

![default_1 preset 2](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-2.png)

Preset 2:

![default_2 preset 1](https://github.com/Promptorium/promptorium/blob/develop/screenshots/screenshot-5.png)


## Prerequisites

- A terminal emulator that supports ANSI escape codes
- A patched [Nerdfont](https://www.nerdfonts.com/) for your terminal

## Installation


### Debian/Ubuntu

You can add the repository and install promptorium using the following commands:

```bash
curl -s http://apt.promptorium.org/gpg-key.public | sudo tee /etc/apt/keyrings/promptorium-gpg.public
echo "deb [arch=amd64 signed-by=/etc/apt/keyrings/promptorium-gpg.public] https://apt.promptorium.org/ unstable main" | sudo tee /etc/apt/sources.list.d/promptorium.list
sudo apt update
sudo apt install promptorium
```

Or you can download the deb package from the [releases page](https://github.com/Promptorium/promptorium/releases) and install it using the following command:

```bash
sudo dpkg -i promptorium_[version]-1_[arch].deb
```

## Usage

Promptorium sets the shell's prompt to the output of a command. You can use the command as a default value in your prompt. For example, if you want to use the default prompt command, you can do the following:
```bash
promptorium prompt --shell bash --exit-code $?
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
		}
	
	]
}
```
