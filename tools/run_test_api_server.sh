#!/usr/bin/env bash
# The script builds and runs test Kentik API server.

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_DIR=$(cd -- "$SCRIPT_DIR" && cd ../ && pwd)

TEST_API_SERVER_ENDPOINT="localhost:9955"

go build -o "$REPO_DIR/test_api_server" github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver
"$REPO_DIR/test_api_server" -addr ${TEST_API_SERVER_ENDPOINT} -synthetics  "$REPO_DIR/synthetics/test-data.json" &
