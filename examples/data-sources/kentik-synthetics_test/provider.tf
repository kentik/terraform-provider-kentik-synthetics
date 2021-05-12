terraform {
  required_providers {
    kentik-synthetics = {
      version = "~> 0.1.0"
      source  = "kentik.com/automation/kentik-synthetics"
    }
  }
}

provider "kentik-synthetics" {
  # email, token and api_url are read from KTAPI_AUTH_EMAIL, KTAPI_AUTH_TOKEN, KTAPI_URL env variables
}
