#!/usr/bin/env bash
# Run the example using live Kentik API server.

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(cd -- "$SCRIPT_DIR" && cd ../../../ && pwd)

function run() {
    check_env

    stage "Build & install plugin"
    make --directory "$REPO_DIR" install || die

    pushd "$SCRIPT_DIR" > /dev/null || die

    stage "Terraform init & apply"
    rm -rf .terraform .terraform.lock.hcl

    # export TF_LOG=ERROR
    terraform init || die
    terraform apply

    popd > /dev/null || die
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

function die() {
    echo "Error. Exit 1"
    exit 1
}

run
