#!/usr/bin/env bash
# Bash utility functions used within the project.

function stage() {
    BOLD_BLUE="\e[1m\e[96m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function warn() {
    RED="\e[31m"
    RESET="\e[0m"
    msg="$1"

    echo -e "$RED$msg$RESET"
}

function pause() {
    read -r -p "Press any key to continue or use Ctrl-C to stop..."
}

function die() {
    echo "Error. Exit 1"
    exit 1
}
