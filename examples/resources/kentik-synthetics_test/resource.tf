resource "kentik-synthetics_test" "example-hostname-test" {
  name   = "example-hostname-test"
  type   = "hostname"
  status = "TEST_STATUS_ACTIVE"
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
      latency_critical_stddev      = 11
      latency_warning_stddev       = 12
      jitter_critical_stddev       = 13
      jitter_warning_stddev        = 14
      http_latency_critical_stddev = 15
      http_latency_warning_stddev  = 16
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
      delay  = 4
    }
    trace {
      period   = 60
      count    = 3
      protocol = "udp"
      port     = 33434
      expiry   = 22500
      limit    = 30
      delay    = 7
    }
    port     = 443
    protocol = "tcp"
    family   = "IP_FAMILY_V4"
    servers = [
      "192.0.2.1",
      "192.0.2.2",
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
  status = "TEST_STATUS_ACTIVE"
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
  status = "TEST_STATUS_ACTIVE"
  settings {
    ip {
      targets = [
        "192.0.2.1",
        "192.0.2.2"
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
  status = "TEST_STATUS_ACTIVE"
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

resource "kentik-synthetics_test" "minimal-dns-test" {
  name   = "minimal-dns-test"
  type   = "dns"
  status = "TEST_STATUS_ACTIVE"
  settings {
    dns {
      target = "www.example.com"
      type   = "DNS_RECORD_CNAME"
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
    period = 60
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
    port     = 443
    protocol = "tcp"
    family   = "IP_FAMILY_V4"
    servers = [
      "192.0.2.1",
      "192.0.2.2",
    ]
    rollup_level = 1
  }
}

resource "kentik-synthetics_test" "minimal-url-test" {
  name   = "minimal-url-test"
  type   = "url"
  status = "TEST_STATUS_ACTIVE"
  settings {
    url {
      target = "https://example.com"
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

resource "kentik-synthetics_test" "minimal-network-grid-test" {
  name   = "minimal-network-grid-test"
  type   = "network_grid"
  status = "TEST_STATUS_ACTIVE"
  settings {
    network_grid {
      targets = ["192.0.2.1"]
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

resource "kentik-synthetics_test" "minimal-page-load-test" {
  name   = "minimal-page-load-test"
  type   = "page_load"
  status = "TEST_STATUS_ACTIVE"
  settings {
    page_load {
      target = "https://example.com"
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
    period = 60
    expiry = 5000
    tasks = [
      "page_load"
    ]
    ping {
      period = 60
    }
    trace {
      period   = 60
      protocol = "udp"
    }
    port     = 443
    protocol = "tcp"
    family   = "IP_FAMILY_V4"
    servers = [
      "192.0.2.1",
      "192.0.2.2",
    ]
    rollup_level = 1
  }
}

resource "kentik-synthetics_test" "minimal-dns-grid-test" {
  name   = "minimal-dns-grid-test"
  type   = "dns_grid"
  status = "TEST_STATUS_ACTIVE"
  settings {
    dns_grid {
      targets = ["www.example.com"]
      type    = "DNS_RECORD_CNAME"
    }
    agent_ids = [
      "817",
      "818",
      "819"
    ]
    period = 60
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
    port     = 443
    protocol = "tcp"
    family   = "IP_FAMILY_V4"
    servers = [
      "192.0.2.1",
      "192.0.2.2",
    ]
    rollup_level = 1
  }
}

resource "kentik-synthetics_test" "minimal-application-mesh-test" {
  name   = "minimal-application-mesh-test"
  type   = "application_mesh"
  status = "TEST_STATUS_ACTIVE"
  settings {
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
    protocol     = "icmp"
    family       = "IP_FAMILY_V4"
    rollup_level = 1
  }
}
