#!/bin/bash

# This script builds the promptorium deb package
# It needs the PROMPTORIUM_VERSION environment variable to be set

architectures=("linux/amd64")
function main() {

    # Exit if no version is set
    if [ -z "$PROMPTORIUM_VERSION" ]; then
        echo "Please set the PROMPTORIUM_VERSION environment variable"
        exit 1
    fi

    for architecture in "${architectures[@]}"
    do
        platform_split=(${architecture//\// })
        os=${platform_split[0]}
        arch=${platform_split[1]}
        if [ "$os" == "linux" ]; then
            echo "Building for $os/$arch"
            create_deb_package $arch
        fi

    done
}

function create_deb_package() {
    local arch=$1

    local deb_directory="./build/deb/promptorium_"$PROMPTORIUM_VERSION"-1_$arch"

    echo "Creating debian package version $PROMPTORIUM_VERSION for $arch"

    # Create directories
    mkdir -p $deb_directory/usr/bin
    mkdir -p $deb_directory/DEBIAN

    # Copy binary
    cp ./build/promptorium_"$PROMPTORIUM_VERSION"_linux_"$arch" $deb_directory/usr/bin/promptorium

    # Create control file
    create_control_file

    copy_config_files

    # Create package
    dpkg --build $deb_directory

}

function copy_config_files() {
    # Copy config files
    mkdir -p $deb_directory/usr/share/promptorium/conf
    cp -r ./conf/* $deb_directory/usr/share/promptorium/conf

}

function create_control_file() {
    echo "Package: promptorium 
Version: $PROMPTORIUM_VERSION
Maintainer: Vladislav Parfeniuc
Homepage: https://github.com/Promptorium/promptorium
Architecture: $arch
Depends: git
Description: A modular and configurable terminal prompt builder" > ./build/deb/promptorium_"$PROMPTORIUM_VERSION"-1_$arch/DEBIAN/control

}

main "$@"
