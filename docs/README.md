
## What is Promptorium?
From a high-level perspective, Promptorium is a modular and configurable terminal prompt builder.

On a more technical level, Promptorium is a command-line tool that generates a prompt based on a configuration file.

It does so not by directly modifying the shell's configuration file, but by setting the shell's prompt to the output of the `promptorium prompt` command.

This means that you can use Promptorium to create prompts for any shell, not just bash or zsh.


## Get Started
To start using promptorium, you need to install it on your system.

> [!NOTE] 
Make sure you are using a patched [Nerdfont](https://www.nerdfonts.com/) for your terminal, as promptorium uses powerline symbols and icons to create its prompt. We recommend using [Fira Code](https://github.com/tonsky/FiraCode) or [JetBrains Mono](https://www.jetbrains.com/lp/mono/) as your Nerdfont.



### Debian/Ubuntu

You can add the repository and install promptorium using the following commands:

```bash
curl -s https://apt.promptorium.org/gpg-key.public | sudo tee /etc/apt/keyrings/promptorium-gpg.public
echo "deb [arch=amd64 signed-by=/etc/apt/keyrings/promptorium-gpg.public] https://apt.promptorium.org/ unstable main" | sudo tee /etc/apt/sources.list.d/promptorium.list
sudo apt update
sudo apt install promptorium
```

Or you can download the deb package from the [releases page](https://github.com/Promptorium/promptorium/releases) and install it using the following command:

```bash
sudo dpkg -i promptorium_[version]-1_[arch].deb
```


Now restart your terminal and you should be good to go!


## Configuration

The configuration files are located at `~/.config/promptorium/`.

### Configuration file
The configuration file is a JSON file that contains an array of components.
Each component has a name, content, and style.

The content is the actual content of the module, while the style is used to customize the appearance of the module.

Here is an example of a configuration file:

```json
{
	"components" : [
		{
			"name" : "user_component",
			"content": {
				"module": "user",
				"icon": ""
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
				"icon": ""
			},
			"style" : {
				"background_color": "$primary",
				"foreground_color": "black",
				"margin" :"1"
			}
		}
    ]
}

```

In this example, there are two components:

- user_component: The user's username
- cwd_component: The current working directory

Each component has a style that is used to customize its appearance.