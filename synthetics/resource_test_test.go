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

//nolint: funlen
func checkResourceTestCreate() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet(testResource, "id"),
		resource.TestCheckResourceAttr(testResource, "name", "dummy-name"),
		resource.TestCheckResourceAttr(testResource, "type", "dummy-type"),
		resource.TestCheckResourceAttr(testResource, "status", "TEST_STATUS_ACTIVE"),
		resource.TestCheckResourceAttr(testResource, "settings.0.hostname.0.target", "dummy-ht"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ip.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ip.0.targets.0", "101.102.103.104"),
		resource.TestCheckResourceAttr(testResource, "settings.0.agent.0.target", "dummy-at"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.target", "dummy-ft"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.target_refresh_interval_millis", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.max_tasks", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.flow.0.type", "dummy-ftt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.site.0.target", "dummy-st"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tag.0.target", "dummy-tt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns.0.target", "dummy-dt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.url.0.target", "dummy-ut"),
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
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_grace_period", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_unit", "m"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_window", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_times", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.0", "nc-1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.1", "nc-2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.period", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.count", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.expiry", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.period", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.count", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.protocol", "udp"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.port", "4"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.expiry", "5"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.limit", "6"),
		resource.TestCheckResourceAttr(testResource, "settings.0.port", "443"),
		resource.TestCheckResourceAttr(testResource, "settings.0.protocol", "icmp"),
		resource.TestCheckResourceAttr(testResource, "settings.0.family", "IP_FAMILY_DUAL"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.#", "2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.0", "server-one"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.1", "server-two"),
		resource.TestCheckResourceAttr(testResource, "settings.0.use_local_ip", "true"),
		resource.TestCheckResourceAttr(testResource, "settings.0.reciprocal", "false"),
		resource.TestCheckResourceAttr(testResource, "settings.0.rollup_level", "1"),
	)
}

//nolint: funlen
func checkResourceTestUpdate() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet(testResource, "id"),
		resource.TestCheckResourceAttr(testResource, "name", "updated-name"),
		resource.TestCheckResourceAttr(testResource, "type", "updated-type"),
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
		resource.TestCheckResourceAttr(testResource, "settings.0.site.0.target", "updated-st"),
		resource.TestCheckResourceAttr(testResource, "settings.0.tag.0.target", "updated-tt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.dns.0.target", "updated-dt"),
		resource.TestCheckResourceAttr(testResource, "settings.0.url.0.target", "updated-ut"),
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
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_grace_period", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_unit", "y"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_time_window", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.activation_times", "30"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.#", "3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.0", "nc-1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.1", "nc-2"),
		resource.TestCheckResourceAttr(testResource, "settings.0.monitoring_settings.0.notification_channels.2", "nc-3"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.period", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.count", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.ping.0.expiry", "30"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.period", "10"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.count", "20"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.protocol", "quick"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.port", "40"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.expiry", "50"),
		resource.TestCheckResourceAttr(testResource, "settings.0.trace.0.limit", "60"),
		resource.TestCheckResourceAttr(testResource, "settings.0.port", "80"),
		resource.TestCheckResourceAttr(testResource, "settings.0.protocol", "pigeon"),
		resource.TestCheckResourceAttr(testResource, "settings.0.family", "IP_FAMILY_V6"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.#", "1"),
		resource.TestCheckResourceAttr(testResource, "settings.0.servers.0", "server-one"),
		resource.TestCheckResourceAttr(testResource, "settings.0.use_local_ip", "false"),
		resource.TestCheckResourceAttr(testResource, "settings.0.reciprocal", "true"),
		resource.TestCheckResourceAttr(testResource, "settings.0.rollup_level", "10"),
	)
}

const (
	testResource = "kentik-synthetics_test.dummy-test"

	resourceTestCreateConfig = `
		provider "kentik-synthetics" {
			debug = true
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
					targets = tolist([
						"101.102.103.104"
					])
				}
				agent {
					target = "dummy-at"
				}
				flow {
					target = "dummy-ft"
					target_refresh_interval_millis = 1
					max_tasks = 2
					type = "dummy-ftt"
				}
				site {
					target = "dummy-st"
				}
				tag {
					target = "dummy-tt"
				}
				dns {
					target = "dummy-dt"
				}
				url {
					target = "dummy-ut"
				}
				agent_ids = tolist([
					"1",
					"2",
					"3"
				])
				period = 1
				count  = 2
				expiry = 3
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
						201,
					])
					dns_valid_codes = tolist([
						21,
						37
					])
				}
				monitoring_settings {
					activation_grace_period = "1"
					activation_time_unit    = "m"
					activation_time_window  = "2"
					activation_times        = "3"
					notification_channels = tolist([
						"nc-1",
						"nc-2",
					])
				}
				ping {
					period = 1
					count  = 2
					expiry = 3
				}
				trace {
					period   = 1
					count    = 2
					protocol = "udp"
					port     = 4
					expiry   = 5
					limit    = 6
				}
				port     = 443
				protocol = "icmp"
				family   = "IP_FAMILY_DUAL"
				servers = tolist([
					"server-one",
					"server-two",
				])
				use_local_ip = true
				reciprocal   = false
				rollup_level = 1
			}
		}
	`

	resourceTestUpdateConfig = `
		provider "kentik-synthetics" {
			debug = true
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
					targets = tolist([
						"101.102.103.104",
						"201.202.203.204"
					])
				}
				agent {
					target = "updated-at"
				}
				flow {
					target = "updated-ft"
					target_refresh_interval_millis = 10
					max_tasks = 20
					type = "updated-ftt"
				}
				site {
					target = "updated-st"
				}
				tag {
					target = "updated-tt"
				}
				dns {
					target = "updated-dt"
				}
				url {
					target = "updated-ut"
				}
				agent_ids = tolist([
					"1",
					"2"
				])
				period = 10
				count  = 20
				expiry = 30
				limit  = 40
				tasks = tolist([
					"ping",
					"traceroute",
					"brew-coffee"
				])
				health_settings {
					latency_critical      = 10
					latency_warning       = 20
					packet_loss_critical  = 30
					packet_loss_warning   = 40
					jitter_critical       = 50
					jitter_warning        = 60
					http_latency_critical = 70
					http_latency_warning  = 80
					http_valid_codes = tolist([
						200,
						201,
						203
					])
					dns_valid_codes = tolist([
						21
					])
				}
				monitoring_settings {
					activation_grace_period = "10"
					activation_time_unit    = "y"
					activation_time_window  = "20"
					activation_times        = "30"
					notification_channels = tolist([
						"nc-1",
						"nc-2",
						"nc-3",
					])
				}
				ping {
					period = 10
					count  = 20
					expiry = 30
				}
				trace {
					period   = 10
					count    = 20
					protocol = "quick"
					port     = 40
					expiry   = 50
					limit    = 60
				}
				port     = 80
				protocol = "pigeon"
				family   = "IP_FAMILY_V6"
				servers = tolist([
					"server-one"
				])
				use_local_ip = false
				reciprocal   = true
				rollup_level = 10
			}
		}
	`

	resourceTestDestroyConfig = `
		provider "kentik-synthetics" {
			debug = true
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
