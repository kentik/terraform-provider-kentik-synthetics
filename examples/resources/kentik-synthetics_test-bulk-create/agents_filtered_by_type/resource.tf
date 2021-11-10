// Create a test with list of agents of given type.
// Inputs: agentType

data "kentik-synthetics_agents" "agents" {}

locals {
  agent_type = "private" // [private/global]
  agents_ids = [for agent in data.kentik-synthetics_agents.agents.items : agent.id if agent.type == local.agent_type]
}

resource "kentik-synthetics_test" "agents-filtered-by-type-test" {
  name   = "agents-filtered-by-type-test"
  type   = "hostname"
  status = "TEST_STATUS_PAUSED"
  settings {
    hostname {
      target = "www.example.test"
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

output "agents-filtered-by-type-test-output" {
  value = kentik-synthetics_test.agents-filtered-by-type-test
}
