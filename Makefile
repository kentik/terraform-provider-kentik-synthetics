HOSTNAME=kentik.com
NAMESPACE=automation
NAME=kentik-synthetics
BINARY=terraform-provider-${NAME}
VERSION=0.1.0
OS_ARCH=linux_amd64
TEST_API_SERVER_ENDPOINT=localhost:9955

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	# build test API server for the provider under test to talk to
	go build -o test_api_server github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver

	# run test API server; serve predefined data from JSON file
	./test_api_server -addr ${TEST_API_SERVER_ENDPOINT} -synthetics synthetics/test-data.json &

	# give the server some warm up time
	sleep 1

	# run tests:
	# - set KTAPI_URL to our test API server URL - otherwise the provider will try to connect to live Kentik server
	# - set KTAPI_AUTH_EMAIL and KTAPI_AUTH_TOKEN to dummy values - they are required by provider, but not actually used by test API server
	# - set no test caching (-count=1) - beside the provider itself, the tests also depend on the used test API server and test data
	KTAPI_URL="http://${TEST_API_SERVER_ENDPOINT}" KTAPI_AUTH_EMAIL="dummy@acme.com" KTAPI_AUTH_TOKEN="dummy" \
		go test ./... $(TESTARGS) -timeout=5m -count=1 || (pkill -f localhost_apiserver && exit 1) # stop server on error

	# finally, stop the server
	pkill -f test_api_server
