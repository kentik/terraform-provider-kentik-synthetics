// Create a test with list of agents - one agent per country.
// Input: list of countries

data "kentik-synthetics_agents" "agents" {}

locals {
  country_list = ["Poland", "United Kingdom", "Netherlands"]
  country_to_ids_map = {for agent in data.kentik-synthetics_agents.agents.items: agent.country => agent.id...
                          if contains(local.country_list, agent.country)}
  agent_ids = [for key, val in local.country_to_ids_map: val[0]]
}

resource "kentik-synthetics_test" "one_agent_per_country-test" {
  name      = "agents-filtered-by-country-test"
  type      = "hostname"
  device_id = "75702"
  status    = "TEST_STATUS_ACTIVE"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = local.agent_ids
    tasks = tolist([
      "ping",
      "traceroute"
    ])
    health_settings {}
    monitoring_settings {
      activation_grace_period = "2"
      activation_time_unit    = "m"
      activation_time_window  = "5"
      activation_times        = "3"
    }
    ping {
      period = 60
      count  = 5
      expiry = 3000
    }
    trace {
      period   = 60
      count    = 3
      protocol = "udp"
      port     = 33434
      expiry   = 22500
      limit    = 30
    }
    port     = 443
    protocol = "icmp"
    family   = "IP_FAMILY_V6"
    rollup_level = 1
  }
}

output "one_agent_per_country-test-output" {
  value = sort(data.kentik-synthetics_agents.agents.items.id)
}