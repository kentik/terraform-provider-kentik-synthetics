#!/usr/bin/env bash
# Check if generated documentation is up-to-date.
# Run from the repository root.

source "tools/utility_functions.sh" || exit 1

function run() {
    stage "Generate docs"
    make docs || die

    stage "Check if docs directory has changed"
    changes=$(git status --porcelain -- docs)
    if [[ -n "${changes}" ]]; then
        echo "Found changes in docs directory:"
        echo "${changes}"

        echo -e "\nDiff for docs directory:"
        git diff -- docs

        echo -e "\nThis means that the documentation is not up-to-date."
        die "Please run 'make docs' and commit changes."
    fi
}

run
