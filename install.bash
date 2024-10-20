#!/bin/bash

# This script installs promptorium

set -e
function main(){

    install_promptorium
    promptorium init
    
    echo "Successfully installed promptorium"
    echo "Please restart your terminal to apply the changes"

}

function install_promptorium() {

    if [[ -n $(promptorium --version 2>/dev/null) ]]; then
        echo "Promptorium is already installed"
        exit 1
    fi

    if is_using_apt; then
        apt_install
        return 0
    fi

    generic_install    

}

function generic_install() {
    if ! is_using_amd64; then
        echo "Promptorium is only available for amd64 architectures"
        exit 1
    fi
    echo "Installing Promptorium..."
    
    local url
    url=$(curl https://api.github.com/repos/Promptorium/promptorium/releases/latest \
    | grep "browser_download_url.*linux_amd64" | cut -d : -f 2,3 | tr -d \" )
    
    if [[ -z $url ]]; then
        echo "Failed to get download URL"
        exit 1
    fi

    echo "Downloading promptorium binary..."
    sudo wget -q -O /usr/local/bin/promptorium "$url"
    sudo chmod +x /usr/local/bin/promptorium

    download_config

}

function download_config() {

    if [[ -d ~/.config/promptorium ]]; then
        echo "Config directory already exists"
        return 0
    fi

    mkdir -p ~/.config/promptorium
    mkdir -p ~/.config/promptorium/shell
    
    local tarball_url
    tarball_url=$(curl -s https://api.github.com/repos/Promptorium/promptorium/releases/latest \
    | grep "tarball_url" | cut -d : -f 2,3 | tr -d \" | tr -d "," | tr -d " ")

    if [[ -z $tarball_url ]]; then
        echo "Failed to get download URL"
        exit 1
    fi
    
    echo "Creating temporary directory..."
    local temp_dir
    temp_dir=$(mktemp -d)

    echo "Downloading config files..."
    wget -q -O "$temp_dir"/promptorium.tar.gz "$tarball_url"
    tar -xzf "$temp_dir"/promptorium.tar.gz -C "$temp_dir"
    cp -r "$temp_dir"/Promptorium*/conf/* ~/.config/promptorium

    echo "Cleaning up..."
    rm -rf "$temp_dir"
}

function apt_install() {
    
    echo "Installing Promptorium using apt..."

    # Add promptorium gpg key if it doesn't exist
    if [[ ! -f /etc/apt/keyrings/promptorium-gpg.public ]]; then
        echo "Adding promptorium gpg key..."
        local gpg_key
        gpg_key=$(curl -s https://apt.promptorium.org/gpg-key.public)
        echo "$gpg_key" | sudo tee /etc/apt/keyrings/promptorium-gpg.public > /dev/null
    else
        echo "promptorium gpg key already exists"
    
    fi

    # Add promptorium apt repository if it doesn't exist
    if [[ ! -f /etc/apt/sources.list.d/promptorium.list ]]; then
        echo "Adding promptorium apt repository..."
        local repository_url
        repository_url="deb [arch=amd64 signed-by=/etc/apt/keyrings/promptorium-gpg.public] https://apt.promptorium.org/ unstable main"

        echo "$repository_url" | sudo tee /etc/apt/sources.list.d/promptorium.list > /dev/null
    else
        echo "promptorium apt repository already exists"
    fi

    # Install promptorium
    echo "Updating apt repositories..."
    sudo apt update
    sudo apt install promptorium -y
}

function is_using_apt() {
    if [[ -n $(command -v apt-get 2>/dev/null) ]]; then
        return 0
    fi
    if [[ -n $(command -v apt 2>/dev/null) ]]; then
        return 0
    fi
    exit 1
}

function is_using_amd64() {
    if [[ "$(uname -m)" == "x86_64" ]]; then
        return 0
    fi
    return 1
}

main