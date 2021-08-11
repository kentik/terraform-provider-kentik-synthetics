// Create a test with list of agents that have a name containing given substring.
// Inputs: name substring

data "kentik-synthetics_agents" "agents" {
  name_substring = "global-agent"
}

locals {
  agents_ids = [for agent in data.kentik-synthetics_agents.agents.items: agent.id]
}

resource "kentik-synthetics_test" "agents-filtered-by-name-test" {
  name      = "agents-filtered-by-name-test"
  type      = "hostname"
  device_id = "75702"
  status    = "TEST_STATUS_ACTIVE"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = local.agents_ids
    tasks = tolist([
      "ping",
      "traceroute"
    ])
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

output "agents-filtered-by-name-test-output" {
  value = kentik-synthetics_test.agents-filtered-by-name-test
}