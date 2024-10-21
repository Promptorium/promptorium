#!/bin/bash

# This script installs promptorium

set -e
function main(){

    if [[ ! -z $(promptorium --version) ]]; then
        echo "promptorium is already installed"
        exit 1
    fi

    echo "Installing promptorium"
    local url=$(curl https://api.github.com/repos/Promptorium/promptorium/releases/latest \
    | grep "browser_download_url.*linux_amd64" | cut -d : -f 2,3 | tr -d \" )
    
    if [[ -z $url ]]; then
        echo "Failed to get download URL"
        exit 1
    fi

    sudo wget -q -O /usr/local/bin/promptorium $url
    sudo chmod +x /usr/local/bin/promptorium

    echo "Successfully installed promptorium"
    echo "Please restart your terminal and you should be ready to use promptorium!"
}

main $@