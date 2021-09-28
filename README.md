# Terraform Provider for Kentik synthetic tests

## Requirements

- [Go](https://golang.org/doc/install) >= 1.16
- [Terraform](https://www.terraform.io/downloads.html) >= 0.13

## Usage

Provider documentation: <https://registry.terraform.io/providers/kentik/kentik-synthetics/latest/docs>  
Usage examples: [examples](./examples)

### Running examples

Go to folder with Terraform `.tf` definition files for Synthetics resources/data sources ([./examples/**](./examples)):

1. Configure provider with parameters:

```terraform
provider "kentik-synthetics" {
  // Authorization email (required). Can also be specified with KTAPI_AUTH_EMAIL environment variable.
  email = "john@acme.com"
  // Authorization token (required). Can also be specified with KTAPI_AUTH_TOKEN environment variable.
  token = "token"
  // Log payloads flag enables verbose debug logs of requests and responses (optional).
  // Can also be specified with KTAPI_LOG_PAYLOADS environment variable.
  log_payloads = true
}
```

or environment variables:

```bash
export KTAPI_AUTH_EMAIL="john@acme.com"
export KTAPI_AUTH_TOKEN="token"
export KTAPI_LOG_PAYLOADS="true"
```

2. Apply Terraform configuration:

```bash
terraform init
terraform apply # creates resources in Kentik platform
```

3. Clean up resources:

```bash
terraform destroy
```

## Development

Anybody who wants to contribute to development is welcome to provide pull requests.

To work on the provider, install tools listed in [requirements section](#requirements).

Optional tools:

- _golangci-lint_: <https://golangci-lint.run/usage/install/#local-installation>

Development steps:
- Build the provider: `make build`
- Build and install the provider locally: `make install`
- Run tests: `make test`
- Run linter: `golangci-lint run`
- Format the code: `./tools/fmt.sh`
- Generate or update documentation: `go generate`

### Test

Tests run the provider against a `test-api-server` that serves data read from `/synthetics/test-data.json`

This allows to:
- avoid the necessity of providing valid API credentials
- avoid creating resources on remote server
- make the test results more reliable

Running `make test` will:
1. Build and run test-api-server which emulates Kentik API v6 by returning static preconfigured responses
2. Run tests (communication with `test-api-server`)
3. Shut down `test-api-server`

### Debug

For debugging use [Delve debugger](https://github.com/go-delve/delve)

```bash
make build
dlv exec ./terraform-provider-kentik-synthetics
r -debug
c
# attach with terraform following the just-printed out instruction in your terminal
```
