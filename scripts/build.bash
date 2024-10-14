#!/bin/bash

# This script builds the promptorium binary
# It needs the PROMPTORIUM_VERSION environment variable to be set

architectures=("linux/amd64")


function build() {
    local os=$1
    local arch=$2
    local output="../build/promptorium_$PROMPTORIUM_VERSION""_$os""_""$arch"
    echo "Building $output"
    GOOS="$os" GOARCH="$arch" go build -o "$output" -ldflags "-X main.Version=$PROMPTORIUM_VERSION"
}

function main() {

    # Exit if no version is set
    if [ -z "$PROMPTORIUM_VERSION" ]; then
        echo "Please set the PROMPTORIUM_VERSION environment variable"
        exit 1
    fi

    echo "Building version $PROMPTORIUM_VERSION"
    for architecture in "${architectures[@]}"
    do
        platform_split=(${architecture//\// })
	    os=${platform_split[0]}
        arch=${platform_split[1]}
        echo "Building for $os/$arch"
        build "$os" "$arch"
    done
}

main