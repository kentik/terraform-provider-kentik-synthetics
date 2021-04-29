#!/usr/bin/env bash
# The script applies the example Terraform configuration.
# The provider uses stub Kentik API server by default.
# Production Kentik API server can be used by passing "production" positional argument to the script.

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(cd -- "$SCRIPT_DIR" && cd ../../../ && pwd)

PRODUCTION_FLAG="production"
TEST_API_SERVER_ENDPOINT=${TEST_API_SERVER_ENDPOINT:-"localhost:9955"}

function run() {
    if [[ $1 == "$PRODUCTION_FLAG" ]]; then
        stage "Running the example using production Kentik API server"
        warn "Warning: Kentik production resources might be modified and paid actions might be invoked"
        pause
    else
        stage "Running the example using test API server"
        echo "Please make sure that test API server is running at $TEST_API_SERVER_ENDPOINT"
        set_test_env
    fi

    check_env

    stage "Build & install plugin"
    make --directory "$REPO_DIR" install || die

    stage "Terraform init & apply"
    pushd "$SCRIPT_DIR" > /dev/null || die
    rm -rf .terraform .terraform.lock.hcl

    terraform init || die
    terraform apply

    popd > /dev/null || die
}

function set_test_env() {
    export KTAPI_URL="http://$TEST_API_SERVER_ENDPOINT"
    export KTAPI_AUTH_EMAIL="dummy@acme.com"
    export KTAPI_AUTH_TOKEN="dummy"
}

function check_env() {
    if [[ -z "$KTAPI_AUTH_EMAIL" ]]; then
        echo "KTAPI_AUTH_EMAIL env variable must be set to Kentik API account email"
        die
    fi

    if [[ -z "$KTAPI_AUTH_TOKEN" ]]; then
        echo "KTAPI_AUTH_TOKEN env variable must be set to Kentik API authorization token"
        die
    fi
}

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
    read -r -p "Press any key to continue..."
}

function die() {
    echo "Error. Exit 1"
    exit 1
}

run "$1"
