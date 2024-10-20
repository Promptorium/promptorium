#!/bin/bash

# This script builds the promptorium deb package
# It needs the PROMPTORIUM_VERSION environment variable to be set

architectures=("linux/amd64")

function create_control_file() {
    echo "Package: promptorium 
Version: $PROMPTORIUM_VERSION
Maintainer: Vladislav Parfeniuc
Homepage: https://github.com/Promptorium/promptorium
Architecture: $arch
Depends: bash, git, curl
Description: A modular and configurable terminal prompt builder" > ./build/deb/promptorium_"$PROMPTORIUM_VERSION"-1_$arch/DEBIAN/control

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

    # Copy config files
    mkdir -p $deb_directory/usr/share/promptorium/conf
    cp -r ./conf/* $deb_directory/usr/share/promptorium/conf
    cp -r ./shell $deb_directory/usr/share/promptorium/conf/shell

    # Copy postinst script
    cp ./scripts/deb/postinst $deb_directory/DEBIAN/postinst
    chmod +x $deb_directory/DEBIAN/postinst

    # Copy postrm script
    cp ./scripts/deb/postrm $deb_directory/DEBIAN/postrm
    chmod +x $deb_directory/DEBIAN/postrm

    # Create package
    dpkg --build $deb_directory

}


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
        echo "Building for $os/$arch"
        if [ "$os" == "linux" ]; then
            create_deb_package $arch
        fi

    done
}

main $@