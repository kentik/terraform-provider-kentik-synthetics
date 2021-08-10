// Create a test with list of agents that have given ASN.
// Inputs: list of asns

data "kentik-synthetics_agents" "agents" {}

locals {
  asnList = [20473, 11111, 12333]

  agents_ids = compact([for agent in data.kentik-synthetics_agents.agents.items:
                          contains(local.asnList, agent.asn) ? agent.id : ""])
}

resource "kentik-synthetics_test" "private-agents-test" {
  name      = "agents-filtered-by-asn-test"
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

output "private-agents-test-output" {
  value = kentik-synthetics_test.private-agents-test
}