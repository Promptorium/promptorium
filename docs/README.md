---
sidebar_position: 1
---

# Introduction


## What is Promptorium?
From a high-level perspective, Promptorium is a modular and configurable terminal prompt builder.

On a more technical level, Promptorium is a ***command-line tool*** that generates a prompt based on a configuration file.
It does so by setting the shell's prompt to the output of a command, without relying on the shell's built-in functionalities. 
This means that promptorium is almost completely **shell agnostic**.

## Get Started
To start using promptorium, you first need to install it on your system.

:::note 
Make sure you are using a patched [Nerdfont](https://www.nerdfonts.com/) for your terminal. We recommend using [Fira Code](https://github.com/tonsky/FiraCode) or [JetBrains Mono](https://www.jetbrains.com/lp/mono/).
:::

### Installation 

For now, two installation methods are available, but more will be added soon.

#### Debian/Ubuntu

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

#### Manual

You can manually download the executable from the [releases page](https://github.com/Promptorium/promptorium/releases/latest) and place it in your PATH.
The following command will download the latest version of promptorium and place it in `/usr/local/bin`:

```bash

 curl https://api.github.com/repos/Promptorium/promptorium/releases/latest | grep "browser_download_url.*linux_amd64" | cut -d : -f 2,3 | tr -d \" | wget -qi -o /usr/local/bin/promptorium

```

### Initialization

After installing promptorium, you can run `promptorium init` to initialize the config and theme files. It is recommended to run this command after installing promptorium for any of the above installation methods.
After that, restart your terminal and you should be ready to use promptorium!

## Configuration

The configuration files are located at `~/.config/promptorium/`.

Configuration of Promptorium is split in two parts: `conf.json` and `theme.json`