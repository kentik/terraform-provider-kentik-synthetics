# Terraform Provider for Kentik synthetic tests

## Release process

Release process for the provider is based on Github repository tags. Every tag with format `v[0-9].[0-9].[0-9]` will trigger automatic build of package and publish it in <https://registry.terraform.io/>.

To build and release package:
1. Make sure that all code that you want to release is in master branch
2. Create tag with format `v[0-9].[0-9].[0-9]` in Github. [Releases](https://github.com/kentik/terraform-provider-kentik-synthetics/releases) -> Draft a new release -> Put tag version, name and description
3. Go to [Github Actions](https://github.com/kentik/terraform-provider-kentik-synthetics/actions) to monitor the build

## Requirements

- [Go](https://golang.org/doc/install) >= 1.16
- [Terraform](https://www.terraform.io/downloads.html) >= 0.13

## Installation

Build and install the provider so that Terraform can use it:

```bash
make install
```

## Usage

The provider can be configured with parameters or environment variables:

```terraform
provider "kentik-synthetics" {
  // Synthetics API server URL. Can also be specified with KTAPI_URL environment variable.
  api_url = "https://synthetics.api.kentik.com"
  // Authorization email (required). Can also be specified with KTAPI_AUTH_EMAIL environment variable.
  email = "dummy@acme.com"
  // Authorization token (required). Can also be specified with KTAPI_AUTH_TOKEN environment variable.
  token = "token"
  // Debug flag enables verbose debug logs of requests and responses (optional).
  // Can also be specified with TF_SYNTHETICS_DEBUG environment variable.
  debug = true
}
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
- Format the code: `./tools/fmt.sh`

### Development state

TODO non-functional:
- release process
- generate documentation
- document the release process, testing, debugging in the README
- document how to run examples
- acceptance tests
