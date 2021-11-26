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

function pause_and_run() {
    YELLOW="\e[33m"
    RESET="\e[0m"

    echo -e "Press any key to run $YELLOW'$*'$RESET or use Ctrl-C to stop..."
    read -r
    "$@"
}

function die() {
    echo "Error - exit 1"
    exit 1
}
