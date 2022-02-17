#!/usr/bin/env bash
# The script applies the example Terraform configuration.
# The provider uses stub Kentik API server by default.
# Production Kentik API server can be used by passing "production" positional argument to the script.

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(cd -- "$SCRIPT_DIR" && cd ../../../ && pwd)

TEST_API_SERVER_ENDPOINT=${TEST_API_SERVER_ENDPOINT:-"localhost:9955"}

source "$REPO_DIR/tools/utility_functions.sh" || exit 1

run_examples "$1"
