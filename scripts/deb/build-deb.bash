#!/bin/bash

architectures=("linux/amd64" "linux/arm64" "linux/386")



function main() {
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
    echo "Package: promptorium 
Version: $PROMPTORIUM_VERSION
Maintainer: Vladislav Parfeniuc
Homepage: https://github.com/Promptorium/promptorium
Architecture: $arch
Depends: bash, git, curl
Description: A modular terminal prompt builder" > ./build/deb/promptorium_"$PROMPTORIUM_VERSION"-1_$arch/DEBIAN/control

    # Copy config files
    mkdir -p $deb_directory/usr/share/promptorium/conf
    cp -r ./conf/* $deb_directory/usr/share/promptorium/conf
    cp -r ./shell $deb_directory/usr/share/promptorium/conf/shell

    # Copy postinst script
    cp ./scripts/deb/postinst $deb_directory/DEBIAN/postinst
    chmod +x $deb_directory/DEBIAN/postinst

    # Create package
    dpkg --build $deb_directory

}

main $@