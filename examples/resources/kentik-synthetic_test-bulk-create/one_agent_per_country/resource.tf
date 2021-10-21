// Create a test with list of agents - one agent per country.
// Input: list of countries

data "kentik-synthetics_agents" "agents" {}

locals {
  country_list = ["PL", "GB", "NL"]
  country_to_ids_map = {for agent in data.kentik-synthetics_agents.agents.items: agent.country => agent.id...
                          if contains(local.country_list, agent.country)}
  agent_ids = [for key, val in local.country_to_ids_map: sort(val)[0]]
}

resource "kentik-synthetics_test" "one_agent_per_country-test" {
  name      = "agents-filtered-by-country-test"
  type      = "hostname"
  status    = "TEST_STATUS_PAUSED"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = local.agent_ids
    tasks = [
      "ping",
      "traceroute"
    ]
    monitoring_settings {
      activation_time_unit    = "m"
      activation_time_window  = "5"
      activation_times        = "3"
    }
    ping {
      period = 60
    }
    trace {
      period   = 60
      protocol = "udp"
    }
    protocol = "tcp"
    family   = "IP_FAMILY_V6"
    rollup_level = 1
  }
}

output "one_agent_per_country-test-output" {
  value = kentik-synthetics_test.one_agent_per_country-test
}