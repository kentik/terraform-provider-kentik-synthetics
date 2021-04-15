# Terraform Provider for Kentik synthetic tests

## Requirements

- [Go](https://golang.org/doc/install) >= 1.16
- [Terraform](https://www.terraform.io/downloads.html) >= 0.13

## Install

Build and install the plugin so that Terraform can find and use it:

```bash
make install
```

## Development

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To install the provider, run `make install`. This will place the compiled provider binary under `~/.terraform.d/plugins/`.

In order to run the full suite of tests, run `make test`.

### Development state

Implemented endpoint:
- /synthetics/v202101beta1/agents GET

TODO endpoints:
- /synthetics/v202101beta1/agents/{agent.id} PATCH, GET, DELETE
- /synthetics/v202101beta1/health/tests POST
- /synthetics/v202101beta1/tests GET, POST
- /synthetics/v202101beta1/tests/{id} GET, DELETE, PATCH
- /synthetics/v202101beta1/tests/{id}/results/trace POST
- /synthetics/v202101beta1/tests/{id}/status PUT

TODO non-functional:
- release process
- generate documentation
- document the release process, testing, debugging in the README
- document how to run examples
- acceptance tests
