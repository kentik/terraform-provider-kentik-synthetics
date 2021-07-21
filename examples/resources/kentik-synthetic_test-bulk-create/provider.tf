terraform {
  required_providers {
    kentik-synthetics = {
      version = ">= 0.1.0"
      source  = "kentik.com/automation/kentik-synthetics"
    }
  }
}

provider "kentik-synthetics" {}
