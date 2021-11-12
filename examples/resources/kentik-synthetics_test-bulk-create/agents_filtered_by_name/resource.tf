// Create a test with list of agents that have a name containing given substring.
// Inputs: name substring

data "kentik-synthetics_agents" "agents" {}

locals {
  name_substring = "Comcast"
  agents_ids = [
    for agent in data.kentik-synthetics_agents.agents.items : agent.id
    if length(regexall(local.name_substring, agent.name)) > 0
  ]
}

resource "kentik-synthetics_test" "agents-filtered-by-name-test" {
  name   = "agents-filtered-by-name-test"
  type   = "hostname"
  status = "TEST_STATUS_ACTIVE"
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

output "agents-filtered-by-name-test-output" {
  value = kentik-synthetics_test.agents-filtered-by-name-test
}
