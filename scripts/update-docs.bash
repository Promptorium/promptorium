#!/bin/bash

# This script updates the documentation files
# It appends the contents of the "docs/changes.md" file to the "docs/CHANGELOG.md" file, then empties the "docs/changes.md" file

function checks() {

    if [ ! -f "docs/CHANGELOG.md" ]; then
        echo "docs/CHANGELOG.md does not exist"
        exit 1
    fi

    if [ ! -f "docs/changes.md" ]; then
        echo "docs/changes.md does not exist"
        exit 1
    fi
    local changes=$(cat "docs/changes.md")
    if [ -z "$changes" ]; then
        echo "docs/changes.md is empty"
        exit 1
    fi
}
function main() {
    local path_to_repo=$1
    if [ -z "$path_to_repo" ]; then
        path_to_repo=$(pwd)
    fi
    cd $path_to_repo

    checks $@

    echo "Appending changes to CHANGELOG.md"
    echo "" >> "docs/CHANGELOG.md"
    echo "" >> "docs/CHANGELOG.md"
    cat "docs/changes.md" >> "docs/CHANGELOG.md"
    echo "Emptying changes file"
    echo "" > "docs/changes.md"
    echo "Done"
}

main $@