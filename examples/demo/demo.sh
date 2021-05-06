#!/usr/bin/env bash
# Showcase kentik-synthetics Terraform provider against production Kentik API.
# Data sources for Agents and Tests are shown.

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(cd -- "$SCRIPT_DIR" && cd ../../ && pwd)

source "$REPO_DIR/tools/utility_functions.sh" || exit 1

function run() {
    pushd "$SCRIPT_DIR" > /dev/null || die

    check_prerequisites
    check_env
    cleanup_tf_files

    stage "Build & install kentik-synthetics Terraform provider"
    pause_and_run make --directory "$REPO_DIR" install || die
    pause

    stage "Initialize Terraform"
    pause_and_run pygmentize ./provider.tf
    pause

    pause_and_run terraform init || die
    pause

    stage "Read synthetics agents and tests"
    pause_and_run pygmentize ./data-sources.tf
    pause

    pause_and_run terraform apply -auto-approve  || die
    pause

    stage "Finished successfully"
    popd > /dev/null || exit
}

function check_prerequisites() {
    if ! pygmentize -V > /dev/null 2>&1; then
        echo "Please install Pygments: https://pygments.org/"
        die
    fi

    if ! curl -V > /dev/null 2>&1; then
        echo "Please install cURL: https://curl.se/"
        die
    fi

    if ! jq -V > /dev/null 2>&1; then
        echo "Please install jq: https://stedolan.github.io/jq/"
        die
    fi
}

function check_env() {
    stage "Check auth env variables"

    if [[ -z "$KTAPI_AUTH_EMAIL" ]]; then
        echo "KTAPI_AUTH_EMAIL env variable must be set to Kentik API account email"
        die
    fi

    if [[ -z "$KTAPI_AUTH_TOKEN" ]]; then
        echo "KTAPI_AUTH_TOKEN env variable must be set to Kentik API authorization token"
        die
    fi

    echo "Print KTAPI_AUTH_EMAIL"
    echo "$KTAPI_AUTH_EMAIL"
    echo "Print KTAPI_AUTH_TOKEN (first 10 chars)"
    echo "${KTAPI_AUTH_TOKEN:0:10}"

    pause
}

function cleanup_tf_files() {
    rm -rf .terraform .terraform.lock.hcl terraform.tfstate
}

run
