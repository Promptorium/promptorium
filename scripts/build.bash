#!/bin/bash

architectures=("linux/amd64")

function main() {
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


function build() {
    local os=$1
    local arch=$2
    local output="../build/promptorium_$PROMPTORIUM_VERSION""_$os""_""$arch"
    echo "Building $output"
    GOOS="$os" GOARCH="$arch" go build -o "$output" -ldflags "-X main.Version=$PROMPTORIUM_VERSION"
}
main