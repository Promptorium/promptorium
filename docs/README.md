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

#### Install script

Promptorium also provides an install script that can be used to install the latest version of Promptorium on your system. The script will download the latest version of Promptorium for amd64 architecture and install it in your PATH.

```bash
 curl https://raw.githubusercontent.com/Promptorium/promptorium/refs/heads/main/install.bash | bash
```

## Configuration

The configuration files are located at `~/.config/promptorium/`.

Configuration of Promptorium is split in two parts: `conf.json` and `theme.json`