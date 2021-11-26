# Terraform Provider for Kentik synthetic tests

## Requirements

- [Go](https://golang.org/doc/install) >= 1.17
- [Terraform](https://www.terraform.io/downloads.html) >= 0.15

## Usage

Detailed user documentation for the provider can be found [here](hhttps://registry.terraform.io/providers/kentik/kentik-synthetics/latest/docs).

## Development

Anybody who wants to contribute to development is welcome to provide pull requests. To work on the provider, install tools listed in [requirements section](#requirements).

Optional tools:
- _golangci-lint_: <https://golangci-lint.run/usage/install/#local-installation>

Development steps:
- Build the provider: `make build`
- Build and install the provider locally: `make install`
- Run tests: `make test`
- Run golangci-lint: `make lint`
- Format the code: `make fmt`
- Generate the documentation: `make docs`
- Check if generated documentation is up-to-date: `make check-docs`

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
