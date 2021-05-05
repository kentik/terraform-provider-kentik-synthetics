#!/usr/bin/env bash
# The script applies the example Terraform configuration.
# The provider uses stub Kentik API server by default.
# Production Kentik API server can be used by passing "production" positional argument to the script.
# TODO(dfurman): deduplicate with run.sh functions in other examples

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(cd -- "$SCRIPT_DIR" && cd ../../../ && pwd)

TEST_API_SERVER_ENDPOINT=${TEST_API_SERVER_ENDPOINT:-"localhost:9955"}

source "$REPO_DIR/tools/utility_functions.sh" || exit 1

function run() {
    if [[ $1 == "production" ]]; then
        stage "Running the example using production Kentik API server"
        warn "Warning: Kentik production resources might be modified"
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

run "$1"
