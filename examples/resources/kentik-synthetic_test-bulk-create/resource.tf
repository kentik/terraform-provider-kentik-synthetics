// Bulk synthetic tests create use cases:
// 1. Customer wants to automatically create synthetic tests from a set of agents
//    to addresses of interfaces classified as “external”.
// 2. Customer wants to automatically create synthetic tests to his applications hosted in AWS
//    from all of his private agents.
// Use case #2 below.

data "kentik-synthetics_agents" "agents" {}

locals {
  agents = [
    for agent in data.kentik-synthetics_agents.agents:
      agent
  ]
  agents_data = local.agents[2]
  private_agents_ids = compact([for agent in local.agents_data : agent.type == "private" ? agent.id : ""])
}

resource "kentik-synthetics_test" "private-agents-test" {
  name      = "private-agents-test"
  type      = "hostname"
  device_id = "75702"
  status    = "TEST_STATUS_ACTIVE"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = local.private_agents_ids
    period = 61
    count  = 2
    expiry = 5001
    limit  = 4
    tasks = tolist([
      "ping",
      "traceroute"
    ])
    health_settings {
      latency_critical      = 1
      latency_warning       = 2
      packet_loss_critical  = 3
      packet_loss_warning   = 4
      jitter_critical       = 5
      jitter_warning        = 6
      http_latency_critical = 7
      http_latency_warning  = 8
      http_valid_codes = tolist([
        200,
        201
      ])
      dns_valid_codes = tolist([
        1,
        2,
        3
      ])
    }
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
    servers = tolist([
      "server-one",
      "server-two",
    ])
    use_local_ip = true
    reciprocal   = false
    rollup_level = 1
  }
}

output "private-agents-test-output" {
  value = kentik-synthetics_test.private-agents-test
}