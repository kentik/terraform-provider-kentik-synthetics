---
layout: ""
page_title: "Provider: Kentik Synthetics"
description: |-
  The Kentik Synthetics provider is used to work with Kentik synthetics tests.
---

# Kentik Synthetics provider

The Kentik Synthetics provider is used to work with Kentik synthetics tests.

## Example usage

Go to folder with Terraform `.tf` definition files for Synthetics resources/data sources ([./examples/**](./examples)):

1. Configure provider with parameters:

```terraform
provider "kentik-synthetics" {
  // Authorization email (required). Can also be specified with KTAPI_AUTH_EMAIL environment variable.
  email = "john@acme.com"
  // Authorization token (required). Can also be specified with KTAPI_AUTH_TOKEN environment variable.
  token = "token"
}
```

or environment variables:

```bash
export KTAPI_AUTH_EMAIL="john@acme.com"
export KTAPI_AUTH_TOKEN="token"
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

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) Authorization email (required). Can also be specified with KTAPI_AUTH_EMAIL environment variable.
- `token` (String, Sensitive) Authorization token (required). Can also be specified with KTAPI_AUTH_TOKEN environment variable.

### Optional

- `api_url` (String) Synthetics API server URL (optional). Can also be specified with KTAPI_URL environment variable.
- `log_payloads` (Boolean) Log payloads flag enables verbose debug logs of requests and responses (optional). Can also be specified with KTAPI_LOG_PAYLOADS environment variable.
- `retry` (Block List, Max: 1) Configuration for API client retry mechanism (see [below for nested schema](#nestedblock--retry))

<a id="nestedblock--retry"></a>
### Nested Schema for `retry`

Optional:

- `max_attempts` (Number) Maximum number of request retry attempts. Minimum valid value: 1 (0 fallbacks to default). Default: 100. Can also be specified with KTAPI_RETRY_MAX_ATTEMPTS environment variable.
- `max_delay` (String) Maximum delay before request retry. Expected Go time duration format, e.g. 1s (see: <https://pkg.go.dev/time#ParseDuration>). Default: 5m (5 minutes). Can also be specified with KTAPI_RETRY_MAX_DELAY environment variable.
- `min_delay` (String) Minimum delay before request retry. Expected Go time duration format, e.g. 1s (see: <https://pkg.go.dev/time#ParseDuration>). Default: 1s (1 second). Can also be specified with KTAPI_RETRY_MIN_DELAY environment variable.
