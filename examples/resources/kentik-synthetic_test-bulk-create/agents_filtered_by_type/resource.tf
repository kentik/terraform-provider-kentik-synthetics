// Create a test with list of agents of given type.
// Inputs: agentType

data "kentik-synthetics_agents" "agents" {}

locals {
  agentType = "global" // [private/global]
  agents_ids = [for agent in data.kentik-synthetics_agents.agents.items: agent.id if agent.type == local.agentType]
}

resource "kentik-synthetics_test" "agents-filtered-by-type-test" {
  name      = "agents-filtered-by-type-test"
  type      = "hostname"
  status    = "TEST_STATUS_ACTIVE"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = local.agents_ids
    tasks = [
      "ping",
      "traceroute"
    ]
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

output "agents-filtered-by-type-test-output" {
  value = kentik-synthetics_test.agents-filtered-by-type-test
}