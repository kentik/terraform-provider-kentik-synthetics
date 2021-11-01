resource "kentik-synthetics_test" "example-hostname-test" {
  name   = "example-hostname-test"
  type   = "hostname"
  status = "TEST_STATUS_PAUSED"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
    period = 60
    count  = 2
    expiry = 5000
    limit  = 4
    tasks = [
      "ping",
      "traceroute"
    ]
    health_settings {
      latency_critical      = 1
      latency_warning       = 2
      packet_loss_critical  = 3
      packet_loss_warning   = 4
      jitter_critical       = 5
      jitter_warning        = 6
      http_latency_critical = 7
      http_latency_warning  = 8
      http_valid_codes = [
        200,
        201
      ]
      dns_valid_codes = [
        1,
        2,
        3
      ]
    }
    #    monitoring_settings {
    #      // Notice: currently "notification_channels" field cannot be manipulated
    #      notification_channels = [
    #      "dummy-channel-1",
    #      "dummy-channel-2",
    #      ]
    #    }
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
    protocol = "tcp"
    family   = "IP_FAMILY_V4"
    servers = [
      "server-one",
      "server-two",
    ]
    rollup_level = 1
  }
}

output "hostname-test" {
  value = kentik-synthetics_test.example-hostname-test
}

resource "kentik-synthetics_test" "minimal-hostname-test" {
  name   = "minimal-hostname-test"
  type   = "hostname"
  status = "TEST_STATUS_PAUSED"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
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

resource "kentik-synthetics_test" "minimal-ip-test" {
  name   = "minimal-ip-test"
  type   = "ip"
  status = "TEST_STATUS_PAUSED"
  settings {
    ip {
      targets = [
        "127.0.0.1",
        "127.0.0.2"
      ]
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
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

resource "kentik-synthetics_test" "minimal-agent-test" {
  name   = "minimal-agent-test"
  type   = "agent"
  status = "TEST_STATUS_PAUSED"
  settings {
    agent {
      target = "1717" # ID of private agent
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
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

resource "kentik-synthetics_test" "minimal-url-test" {
  name   = "minimal-url-test"
  type   = "url"
  status = "TEST_STATUS_PAUSED"
  settings {
    url {
      target = "https://dummy.url"
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
    period = 60
    expiry = 5000
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
