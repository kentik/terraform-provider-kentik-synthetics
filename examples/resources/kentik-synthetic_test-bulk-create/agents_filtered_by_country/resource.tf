// Create a test with list of agents that are located in a country from given list.
// Input: list of countries

data "kentik-synthetics_agents" "agents" {}

locals {
  country_list = ["PL", "KH", "TR"]
  agents_ids = [
    for agent in data.kentik-synthetics_agents.agents.items : agent.id
    if contains(local.country_list, agent.country)
  ]
}

resource "kentik-synthetics_test" "agents-filtered-by-country-test" {
  name   = "agents-filtered-by-country-test"
  type   = "hostname"
  status = "TEST_STATUS_PAUSED"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = local.agents_ids
    tasks = [
      "ping",
      "traceroute"
    ]
    ping {
      period = 60
    }
    trace {
      period   = 60
      protocol = "udp"
    }
    port         = 443
    protocol     = "tcp"
    family       = "IP_FAMILY_V4"
    rollup_level = 1
  }
}

output "agents-filtered-by-country-output" {
  value = kentik-synthetics_test.agents-filtered-by-country-test
}
