#!/usr/bin/env bash
# The script stops test Kentik API server (spawned with start_test_api_server.sh script).

pkill -f ./test-api-server
