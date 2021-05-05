#!/usr/bin/env bash
# The script stops test Kentik API server (spawned with run_test_api_server.sh script).

pkill --full ./test-api-server
