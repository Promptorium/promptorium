# Promptorium

> A modular and configurable terminal prompt builder

## What is Promptorium?
Promptorium is a modular and configurable terminal prompt builder. It is designed to be easy to use and configure, while still providing a lot of flexibility and customization options.

## How does it work?
Promptorium is built around a simple concept: components.

You can think of a component as a small, configurable unit of functionality within the prompt. For example, you might have a component for displaying the current directory, a component for displaying the current time, and a component for displaying the current git branch.

You can define components in a configuration file, putting them together to create a prompt.


### Components
Let's take a look at how a component is defined in a configuration file:

```json
{
    "name": "current-directory",
    "content":{
        "module":"cwd",
        "icon":"📂",
    },
    "style": {
        "background_color": "green",
        "color": "white",
        "padding": "1 0",
        "margin": "0 1",
        "start_divider": "",
        "end_divider": "",
    }
}
```



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

Promptorium adds a line to your shell's configuration file (e.g. `.bashrc` or `.zshrc`) that sources the promptorium script.
This script sets the shell's prompt to the output of the `promptorium prompt` command.
