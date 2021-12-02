HOSTNAME = kentik
NAMESPACE = automation
NAME = kentik-synthetics
BINARY = terraform-provider-${NAME}
VERSION := $(shell echo `git tag --list 'v*' | tail -1 | cut -d v -f 2` | sed -e 's/^$$/0.1.0/')
OS_ARCH := $(shell printf "%s_%s" `go env GOHOSTOS` `go env GOHOSTARCH`)
TEST_API_SERVER_ENDPOINT=localhost:9955

default: install

build:
	go build -o ${BINARY}

check-docs:
	./tools/check_docs.sh

docs:
	go generate

fmt:
	./tools/fmt.sh

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

lint:
	golangci-lint run

test:
	TEST_API_SERVER_ENDPOINT=${TEST_API_SERVER_ENDPOINT} ./tools/start_test_api_server.sh

	# run tests:
	# - set KTAPI_URL to our test API server URL - otherwise the provider will try to connect to live Kentik server
	# - set KTAPI_AUTH_EMAIL and KTAPI_AUTH_TOKEN to dummy values - they are required by provider,
	#   but not actually used by test API server
	# - set no test caching (-count=1) - beside the provider itself, the tests also depend on the
	#   test API server and test data
	KTAPI_URL="http://${TEST_API_SERVER_ENDPOINT}" KTAPI_AUTH_EMAIL="dummy@acme.com" KTAPI_AUTH_TOKEN="dummy" \
		go test ./... $(TESTARGS) -timeout=5m -count=1 || (./tools/stop_test_api_server.sh && exit 1)

	 ./tools/stop_test_api_server.sh

.PHONY: build check-docs docs fmt install lint test
