#!/usr/bin/env bash
# Check if generated documentation is up-to-date.
# Run from the repository root.

source "tools/utility_functions.sh" || exit 1

function run() {
    stage "Generate docs"
    go generate || die

    stage "Check if docs directory has changed"
    if [[ -n $(git status --porcelain -- docs) ]]; then
        echo "Found changes in docs directory:"
        git status --porcelain -- docs

        echo -e "\nDiff for docs directory:"
        git diff -- docs

        echo -e "\nThis means that the documentation is not up-to-date."
        echo "Please run 'make docs' and commit changes."
        die
    fi
}

run
