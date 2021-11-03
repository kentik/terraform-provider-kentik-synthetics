#!/usr/bin/env bash
# The script builds and runs test Kentik API server.
# It serves predefined data from JSON file.

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(dirname "${SCRIPT_DIR}")

TEST_API_SERVER_ENDPOINT=${TEST_API_SERVER_ENDPOINT:-"localhost:9955"}

go build -o "$REPO_DIR/test-api-server" github.com/kentik/community_sdk_golang/kentikapi/fakeapiserver || exit 1
"$REPO_DIR/test-api-server" -addr "${TEST_API_SERVER_ENDPOINT}" -synthetics  "$REPO_DIR/synthetics/test-data.json" &
