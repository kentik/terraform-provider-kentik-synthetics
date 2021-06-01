resource "kentik-synthetics_test" "example-hostname-test" {
  name      = "example-hostname-test"
  type      = "hostname"
  device_id = "75702"
  status    = "TEST_STATUS_ACTIVE"
  settings {
    hostname {
      target = "www.example.com"
    }
    agent_ids = tolist([
      "817",
      "818",
      "819"
    ])
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
      // Notice: currently "notification_channels" field cannot be manipulated
      //      notification_channels = tolist([
      //        "dummy-channel-1",
      //        "dummy-channel-2",
      //      ])
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
    // Notice: specific values expected
    // TODO(dfurman): provide allowed values
    // target_type  = ""
    // target_value = ""
    use_local_ip = true
    reciprocal   = false
    rollup_level = 1
  }
}

output "hostname-test" {
  value = kentik-synthetics_test.example-hostname-test
}
