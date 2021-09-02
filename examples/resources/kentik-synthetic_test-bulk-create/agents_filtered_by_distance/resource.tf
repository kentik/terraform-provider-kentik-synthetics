// Create a test with list of agents that are located within given distance constraints from coordinates.
// Inputs: latitude, longitude, min_distance, max_distance.

data "kentik-synthetics_agents" "agents" {
  latitude  = 50.55
  longitude = 5.5
  min_distance = 1000
  max_distance = 9000
}

locals {
  agents_ids = [for agent in data.kentik-synthetics_agents.agents.items: agent.id]
}

resource "kentik-synthetics_test" "agents-within-distance-test" {
  name      = "agents-within-radius-test"
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

output "agents-within-distance-test-output" {
  value = kentik-synthetics_test.agents-within-distance-test
}