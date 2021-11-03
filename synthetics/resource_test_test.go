package synthetics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceTest(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config: resourceTestCreateConfig,
			Check:  checkResourceTestCreate(),
		}, {
			Config: resourceTestUpdateConfig,
			Check:  checkResourceTestUpdate(),
		}, {
			Config: resourceTestDestroyConfig,
			Check: resource.ComposeTestCheckFunc(
				checkResourceDoesNotExist(testResource),
			),
		}},
	})
}

func checkResourceTestCreate() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet(testResource, "id"),
		resource.TestCheckResourceAttr(testResource, "name", "dummy-name"),
		resource.TestCheckResourceAttr(testResource, "type", "dummy-type"),
		resource.TestCheckNoResourceAttr(testResource, "device_id"),
		resource.TestCheckResourceAttr(testResource, "status", "TEST_STATUS_ACTIVE"),
		resource.TestCheckResourceAttr(testResource, "settings.0.hostname.0.target", "dummy-ht"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ip.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ip.0.targets.0", "101.102.103.104"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent.0.target", "dummy-at"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.target", "dummy-ft"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.target_refresh_interval_millis", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.max_tasks", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.type", "dummy-ftt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.inet_direction", "dummy-id"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.direction", "dummy-d"),
		resource.TestCheckResourceAttr(testResource, "settings.0.site.0.target", "dummy-st"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tag.0.target", "dummy-tt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns.0.target", "dummy-dt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns.0.type", "DNS_RECORD_CNAME"),
		resource.TestCheckResourceAttr(testResource, "settings.0.url.0.target", "dummy-ut"),
		resource.TestCheckResourceAttr(testResource, "settings.0.network_grid.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.network_grid.0.targets.0", "dummy-ng-target"),
		resource.TestCheckResourceAttr(testResource, "settings.0.page_load.0.target", "dummy-pl-target"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns_grid.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns_grid.0.targets.0", "dummy-dg-target"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns_grid.0.type", "DNS_RECORD_A"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.application_mesh"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent_ids.#", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent_ids.0", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent_ids.1", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent_ids.2", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.period", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.count", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.expiry", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.limit", "4"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tasks.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tasks.0", "ping"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tasks.1", "traceroute"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_critical", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_warning", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.packet_loss_critical", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.packet_loss_warning", "4"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_critical", "5"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_warning", "6"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_critical", "7"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_warning", "8"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_valid_codes.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_valid_codes.0", "200"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_valid_codes.1", "201"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.dns_valid_codes.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.dns_valid_codes.0", "21"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.dns_valid_codes.1", "37"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_critical_stddev", "11"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_warning_stddev", "12"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_critical_stddev", "13"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_warning_stddev", "14"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_critical_stddev", "15"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_warning_stddev", "16"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_grace_period"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_unit"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_window"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_times"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.0", "nc-1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.1", "nc-2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.period", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.count", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.expiry", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.delay", "4"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.period", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.count", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.protocol", "udp"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.port", "4"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.expiry", "5"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.limit", "6"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.delay", "7"),
		resource.TestCheckResourceAttr(testResource, "settings.0.port", "443"),
		resource.TestCheckResourceAttr(testResource, "settings.0.protocol", "icmp"),
		resource.TestCheckResourceAttr(testResource, "settings.0.family", "IP_FAMILY_DUAL"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.0", "server-one"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.1", "server-two"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.use_local_ip"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.reciprocal"),
		resource.TestCheckResourceAttr(testResource, "settings.0.rollup_level", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.period", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.expiry", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.method", "GET"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.headers.dummy-header-key", "dummy-header-value"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.body", "dummy-body"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.ignore_tls_errors", "true"),
		resource.TestCheckResourceAttr(
			testResource, "settings.0.http.0.css_selectors.dummy-selector-key", "dummy-selector-value",
		),
	)
}

func checkResourceTestUpdate() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet(testResource, "id"),
		resource.TestCheckResourceAttr(testResource, "name", "updated-name"),
		resource.TestCheckResourceAttr(testResource, "type", "updated-type"),
		resource.TestCheckNoResourceAttr(testResource, "device_id"),
		resource.TestCheckResourceAttr(testResource, "status", "TEST_STATUS_PAUSED"),
		resource.TestCheckResourceAttr(testResource, "settings.0.hostname.0.target", "updated-ht"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ip.0.targets.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ip.0.targets.0", "101.102.103.104"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ip.0.targets.1", "201.202.203.204"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent.0.target", "updated-at"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.target", "updated-ft"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.target_refresh_interval_millis", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.max_tasks", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.type", "updated-ftt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.inet_direction", "updated-id"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.direction", "updated-d"),
		resource.TestCheckResourceAttr(testResource, "settings.0.site.0.target", "updated-st"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tag.0.target", "updated-tt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns.0.target", "updated-dt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns.0.type", "DNS_RECORD_DNAME"),
		resource.TestCheckResourceAttr(testResource, "settings.0.url.0.target", "updated-ut"),
		resource.TestCheckResourceAttr(testResource, "settings.0.network_grid.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.network_grid.0.targets.0", "updated-ng-target"),
		resource.TestCheckResourceAttr(testResource, "settings.0.page_load.0.target", "updated-pl-target"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns_grid.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns_grid.0.targets.0", "updated-dg-target"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns_grid.0.type", "DNS_RECORD_AAAA"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.application_mesh"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent_ids.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent_ids.0", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent_ids.1", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.period", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.count", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.expiry", "30"),
		resource.TestCheckResourceAttr(testResource, "settings.0.limit", "40"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tasks.#", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tasks.0", "ping"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tasks.1", "traceroute"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tasks.2", "brew-coffee"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_critical", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_warning", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.packet_loss_critical", "30"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.packet_loss_warning", "40"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_critical", "50"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_warning", "60"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_critical", "70"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_warning", "80"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_valid_codes.#", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_valid_codes.0", "200"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_valid_codes.1", "201"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_valid_codes.2", "203"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.dns_valid_codes.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.dns_valid_codes.0", "21"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_critical_stddev", "110"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.latency_warning_stddev", "120"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_critical_stddev", "130"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.jitter_warning_stddev", "140"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_critical_stddev", "150"),
		resource.TestCheckResourceAttr(testResource, "settings.0.health_settings.0.http_latency_warning_stddev", "160"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_grace_period"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_unit"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_window"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_times"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.#", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.0", "nc-1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.1", "nc-2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.2", "nc-3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.period", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.count", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.expiry", "30"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.delay", "40"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.period", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.count", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.protocol", "quick"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.port", "40"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.expiry", "50"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.limit", "60"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.delay", "70"),
		resource.TestCheckResourceAttr(testResource, "settings.0.port", "80"),
		resource.TestCheckResourceAttr(testResource, "settings.0.protocol", "pigeon"),
		resource.TestCheckResourceAttr(testResource, "settings.0.family", "IP_FAMILY_V6"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.0", "server-one"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.use_local_ip"),
		resource.TestCheckNoResourceAttr(testResource, "settings.0.reciprocal"),
		resource.TestCheckResourceAttr(testResource, "settings.0.rollup_level", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.period", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.expiry", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.method", "POST"),
		resource.TestCheckResourceAttr(
			testResource, "settings.0.http.0.headers.dummy-header-key", "updated-header-value",
		),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.body", "updated-body"),
		resource.TestCheckResourceAttr(testResource, "settings.0.http.0.ignore_tls_errors", "false"),
		resource.TestCheckResourceAttr(
			testResource, "settings.0.http.0.css_selectors.dummy-selector-key", "updated-selector-value",
		),
	)
}

const (
	testResource = "kentik-synthetics_test.dummy-test"

	resourceTestCreateConfig = `
		provider "kentik-synthetics" {
			log_payloads = true
		}

		resource "kentik-synthetics_test" "dummy-test" {
			name      = "dummy-name"
			type      = "dummy-type"
			status    = "TEST_STATUS_ACTIVE"
			settings {
				hostname {
					target = "dummy-ht"
				}
				ip {
					targets = [
						"101.102.103.104"
					]
				}
				agent {
					target = "dummy-at"
				}
				flow {
					target = "dummy-ft"
					target_refresh_interval_millis = 1
					max_tasks = 2
					type = "dummy-ftt"
					inet_direction = "dummy-id"
					direction = "dummy-d"
				}
				site {
					target = "dummy-st"
				}
				tag {
					target = "dummy-tt"
				}
				dns {
					target = "dummy-dt"
					type = "DNS_RECORD_CNAME"
				}
				url {
					target = "dummy-ut"
				}
				network_grid {
					targets = ["dummy-ng-target"]
				}
				page_load {
					target = "dummy-pl-target"
				}
				dns_grid {
					targets = ["dummy-dg-target"]
					type = "DNS_RECORD_A"
				}
				agent_ids = [
					"1",
					"2",
					"3"
				]
				period = 1
				count  = 2
				expiry = 3
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
						201,
					]
					dns_valid_codes = [
						21,
						37
					]
					latency_critical_stddev      = 11
					latency_warning_stddev       = 12
					jitter_critical_stddev       = 13
					jitter_warning_stddev        = 14
					http_latency_critical_stddev = 15
					http_latency_warning_stddev  = 16
				}
				monitoring_settings {
					notification_channels = [
						"nc-1",
						"nc-2",
					]
				}
				ping {
					period = 1
					count  = 2
					expiry = 3
					delay = 4
				}
				trace {
					period   = 1
					count    = 2
					protocol = "udp"
					port     = 4
					expiry   = 5
					limit    = 6
					delay    = 7
				}
				port     = 443
				protocol = "icmp"
				family   = "IP_FAMILY_DUAL"
				servers = [
					"server-one",
					"server-two",
				]
				rollup_level = 1
				http {
				  period = 1
				  expiry = 2
				  method = "GET"
				  headers = {
					dummy-header-key = "dummy-header-value"
				  }
				  body = "dummy-body"
				  ignore_tls_errors = true
				  css_selectors = {
					dummy-selector-key = "dummy-selector-value"
				  }
				}
			}
		}
	`

	resourceTestUpdateConfig = `
		provider "kentik-synthetics" {
			log_payloads = true
		}

		resource "kentik-synthetics_test" "dummy-test" {
			name      = "updated-name"
			type      = "updated-type"
			status    = "TEST_STATUS_PAUSED"
			settings {
				hostname {
					target = "updated-ht"
				}
				ip {
					targets = [
						"101.102.103.104",
						"201.202.203.204"
					]
				}
				agent {
					target = "updated-at"
				}
				flow {
					target = "updated-ft"
					target_refresh_interval_millis = 10
					max_tasks = 20
					type = "updated-ftt"
					inet_direction = "updated-id"
					direction = "updated-d"
				}
				site {
					target = "updated-st"
				}
				tag {
					target = "updated-tt"
				}
				dns {
					target = "updated-dt"
					type = "DNS_RECORD_DNAME"
				}
				url {
					target = "updated-ut"
				}
				network_grid {
					targets = ["updated-ng-target"]
				}
				page_load {
					target = "updated-pl-target"
				}
				dns_grid {
					targets = ["updated-dg-target"]
					type = "DNS_RECORD_AAAA"
				}
				agent_ids = [
					"1",
					"2"
				]
				period = 10
				count  = 20
				expiry = 30
				limit  = 40
				tasks = [
					"ping",
					"traceroute",
					"brew-coffee"
				]
				health_settings {
					latency_critical      = 10
					latency_warning       = 20
					packet_loss_critical  = 30
					packet_loss_warning   = 40
					jitter_critical       = 50
					jitter_warning        = 60
					http_latency_critical = 70
					http_latency_warning  = 80
					http_valid_codes = [
						200,
						201,
						203
					]
					dns_valid_codes = [
						21
					]
					latency_critical_stddev      = 110
					latency_warning_stddev       = 120
					jitter_critical_stddev       = 130
					jitter_warning_stddev        = 140
					http_latency_critical_stddev = 150
					http_latency_warning_stddev  = 160
				}
				monitoring_settings {
					notification_channels = [
						"nc-1",
						"nc-2",
						"nc-3",
					]
				}
				ping {
					period = 10
					count  = 20
					expiry = 30
					delay = 40
				}
				trace {
					period   = 10
					count    = 20
					protocol = "quick"
					port     = 40
					expiry   = 50
					limit    = 60
					delay    = 70
				}
				port     = 80
				protocol = "pigeon"
				family   = "IP_FAMILY_V6"
				servers = [
					"server-one"
				]
				rollup_level = 10
				http {
				  period = 10
				  expiry = 20
				  method = "POST"
				  headers = {
					dummy-header-key = "updated-header-value"
				  }
				  body = "updated-body"
				  ignore_tls_errors = false
				  css_selectors = {
					dummy-selector-key = "updated-selector-value"
				  }
				}
			}
		}
	`

	resourceTestDestroyConfig = `
		provider "kentik-synthetics" {
			log_payloads = true
		}
	`
)

func checkResourceDoesNotExist(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		_, exists := s.RootModule().Resources[name]
		if exists {
			return fmt.Errorf("resource %q found when not expected", name)
		}

		return nil
	}
}
