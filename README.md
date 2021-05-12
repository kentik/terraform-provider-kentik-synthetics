# Terraform Provider for Kentik synthetic tests

## Requirements

- [Go](https://golang.org/doc/install) >= 1.16
- [Terraform](https://www.terraform.io/downloads.html) >= 0.13

## Installation

Build and install the provider so that Terraform can use it:

```bash
make install
```

## Development

Anybody who wants to contribute to development is welcome to provide pull requests.

To work on the provider, install tools listed in [requirements section](#requirements).

Optional tools:
- _golangci-lint_: [local installation](https://golangci-lint.run/usage/install/#local-installation)

Development steps:
- Build the provider: `make build`
- Build and install the provider locally: `make install`
- Run tests: `make test`
- Run linter: `golangci-lint run`
- Format the code: `go fmt ./...`

### Development state

Implemented endpoint:
- /synthetics/v202101beta1/agents GET
- /synthetics/v202101beta1/agents/{agent.id} GET
- /synthetics/v202101beta1/tests GET
  /synthetics/v202101beta1/tests/{id} GET

TODO endpoints:
- /synthetics/v202101beta1/health/tests POST
- /synthetics/v202101beta1/tests POST
- /synthetics/v202101beta1/tests/{id} DELETE, PATCH
- /synthetics/v202101beta1/tests/{id}/results/trace POST
- /synthetics/v202101beta1/tests/{id}/status PUT

TODO non-functional:
- release process
- generate documentation
- document the release process, testing, debugging in the README
- document how to run examples
- acceptance tests
