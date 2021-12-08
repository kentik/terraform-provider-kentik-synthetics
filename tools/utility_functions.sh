#!/usr/bin/env bash
# Bash utility functions used within the project.

function stage() {
    echo
    colored_echo BOLD_BLUE "$1"
}

function warn() {
    colored_echo RED "$1"
}

function pause() {
    read -r -p "Press any key to continue or use Ctrl-C to stop..."
}

function pause_and_run() {
    echo "Press any key to run '$*' or use Ctrl-C to stop..."
    read -r
    "$@"
}

function die() {
    colored_echo RED "$*" 1>&2
    exit 1
}

function colored_echo() {
    local color="$1"
    local msg="$2"

    if [[ ${BASH_VERSINFO[0]} -lt 4 ]]; then
        # Bash 3.2 distributed with macOS does not support control sequences in echo -e
        color_code=""
    else
        # shellcheck disable=SC2034
        {
            BOLD_BLUE="\e[1m\e[96m"
            RED="\e[31m"
            YELLOW="\e[33m"
        }
        RESET="\e[0m"
        color_code=$(eval echo "\$$color")
    fi

    if [ -n "${color_code}" ]; then
        echo -e "${color_code}${2}${RESET}"
    else
        echo "$msg"
    fi
}
