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

function run_examples() {
    if [[ $1 == "production" ]]; then
        stage "Running the example using production Kentik API server"
        warn "Warning: Kentik production resources might be modified"
        pause
    else
        stage "Running the example using test API server"
        echo "Please make sure that test API server is running at $TEST_API_SERVER_ENDPOINT"
        set_test_env_examples
    fi

    check_env_examples

    stage "Build & install plugin"
    make --directory "$REPO_DIR" install || die

    stage "Terraform init & apply"
    pushd "$SCRIPT_DIR" > /dev/null || die
    rm -rf .terraform .terraform.lock.hcl

    terraform init || die
    terraform apply

    popd > /dev/null || die
}

function set_test_env_examples() {
    export KTAPI_URL="http://$TEST_API_SERVER_ENDPOINT"
    export KTAPI_AUTH_EMAIL="dummy@acme.com"
    export KTAPI_AUTH_TOKEN="dummy"
}

function check_env_examples() {
    if [[ -z "$KTAPI_AUTH_EMAIL" ]]; then
        die "KTAPI_AUTH_EMAIL env variable must be set to Kentik API account email"
    fi

    if [[ -z "$KTAPI_AUTH_TOKEN" ]]; then
        die "KTAPI_AUTH_TOKEN env variable must be set to Kentik API authorization token"
    fi
}
