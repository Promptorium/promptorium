#!/bin/bash

# This script installs promptorium

set -e
function main(){

    if [[ ! -z $(promptorium --version) ]]; then
        echo "Promptorium is already installed"
        exit 1
    fi

    echo "Installing Promptorium..."
    local url=$(curl https://api.github.com/repos/Promptorium/promptorium/releases/latest \
    | grep "browser_download_url.*linux_amd64" | cut -d : -f 2,3 | tr -d \" )
    
    if [[ -z $url ]]; then
        echo "Failed to get download URL"
        exit 1
    fi

    echo "Downloading promptorium binary..."
    sudo wget -q -O /usr/local/bin/promptorium $url
    sudo chmod +x /usr/local/bin/promptorium

    download_config

    echo "Successfully installed promptorium"
    echo "Please restart your terminal and you should be ready to use promptorium!"
}

function download_config() {

    if [[ -d ~/.config/promptorium ]]; then
        echo "Config directory already exists"
        #return 0
    fi

    mkdir -p ~/.config/promptorium
    mkdir -p ~/.config/promptorium/shell
 
    local tarball_url=$(curl -s https://api.github.com/repos/Promptorium/promptorium/releases/latest \
    | grep "tarball_url" | cut -d : -f 2,3 | tr -d \" | tr -d "," | tr -d " ")

    if [[ -z $tarball_url ]]; then
        echo "Failed to get download URL"
        exit 1
    fi
    
    echo "Creating temporary directory..."
    local temp_dir=$(mktemp -d)

    echo "Downloading config files..."
    wget -q -O $temp_dir/promptorium.tar.gz $tarball_url
    tar -xzf $temp_dir/promptorium.tar.gz -C $temp_dir
    cp -r $temp_dir/Promptorium*/conf/* ~/.config/promptorium
    cp -r $temp_dir/Promptorium*/shell/* ~/.config/promptorium/shell

    echo "Cleaning up..."
    rm -rf $temp_dir
}

download_config
#main $@