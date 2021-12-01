#!/usr/bin/env bash
# Check Go module consistency.
# Run from the repository root.

source "tools/utility_functions.sh" || exit 1

function run() {
    stage "Run 'go mod tidy'"
    go mod tidy || die

    stage "Check if go.mod or go.sum file has changed"
    if [[ -n $(git status --porcelain -- go.mod go.sum) ]]; then
        echo "Found changes in go.mod or go.sum file:"
        git diff -- go.mod go.sum

        echo -e "\nThis means that Go module definition is not consistent."
        echo "Please run 'go mod tidy' and commit changes."
        die
    fi
}

run
