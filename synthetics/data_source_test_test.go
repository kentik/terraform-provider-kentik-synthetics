package synthetics_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Note: values checked in tests below are provided by stub API Server from test-data.json (running in background)

func TestDataSourceTest(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config: dataSourceTestConfig,
			Check:  checkDataSourceTest(),
		}},
	})
}

func checkDataSourceTest() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(testDS, "id", "1"),
		resource.TestCheckResourceAttr(testDS, "name", "github.com ip test"),
		resource.TestCheckResourceAttr(testDS, "type", "ip"),
		resource.TestCheckNoResourceAttr(testDS, "device_id"),
		resource.TestCheckResourceAttr(testDS, "status", "TEST_STATUS_ACTIVE"),
		resource.TestCheckResourceAttr(testDS, "settings.0.hostname.0.target", "dummy-ht"),
		resource.TestCheckResourceAttr(testDS, "settings.0.ip.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.ip.0.targets.0", "140.82.121.3"),
		resource.TestCheckResourceAttr(testDS, "settings.0.agent.0.target", "dummy-at"),
		resource.TestCheckResourceAttr(testDS, "settings.0.flow.0.target", "dummy-ft"),
		resource.TestCheckResourceAttr(testDS, "settings.0.flow.0.target_refresh_interval_millis", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.flow.0.max_tasks", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.flow.0.type", "dummy-ftt"),
		resource.TestCheckResourceAttr(testDS, "settings.0.flow.0.inet_direction", "dummy-id"),
		resource.TestCheckResourceAttr(testDS, "settings.0.flow.0.direction", "dummy-d"),
		resource.TestCheckResourceAttr(testDS, "settings.0.site.0.target", "dummy-st"),
		resource.TestCheckResourceAttr(testDS, "settings.0.tag.0.target", "dummy-tt"),
		resource.TestCheckResourceAttr(testDS, "settings.0.dns.0.target", "dummy-dt"),
		resource.TestCheckResourceAttr(testDS, "settings.0.dns.0.type", "DNS_RECORD_CNAME"),
		resource.TestCheckResourceAttr(testDS, "settings.0.url.0.target", "dummy-ut"),
		resource.TestCheckResourceAttr(testDS, "settings.0.network_grid.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.network_grid.0.targets.0", "dummy-ng-target"),
		resource.TestCheckResourceAttr(testDS, "settings.0.page_load.0.target", "dummy-pl-target"),
		resource.TestCheckResourceAttr(testDS, "settings.0.dns_grid.0.targets.#", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.dns_grid.0.targets.0", "dummy-dg-target"),
		resource.TestCheckResourceAttr(testDS, "settings.0.dns_grid.0.type", "DNS_RECORD_A"),
		resource.TestCheckNoResourceAttr(testDS, "settings.0.application_mesh"),
		resource.TestCheckResourceAttr(testDS, "settings.0.agent_ids.#", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.agent_ids.0", "817"),
		resource.TestCheckResourceAttr(testDS, "settings.0.period", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.count", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.expiry", "3"),
		resource.TestCheckResourceAttr(testDS, "settings.0.limit", "4"),
		resource.TestCheckResourceAttr(testDS, "settings.0.tasks.#", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.tasks.0", "ping"),
		resource.TestCheckResourceAttr(testDS, "settings.0.tasks.1", "traceroute"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.latency_critical", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.latency_warning", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.packet_loss_critical", "3"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.packet_loss_warning", "4"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.jitter_critical", "5"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.jitter_warning", "6"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.http_latency_critical", "7"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.http_latency_warning", "8"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.http_valid_codes.#", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.http_valid_codes.0", "200"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.http_valid_codes.1", "201"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.dns_valid_codes.#", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.dns_valid_codes.0", "21"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.dns_valid_codes.1", "37"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.latency_critical_stddev", "11"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.latency_warning_stddev", "12"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.jitter_critical_stddev", "13"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.jitter_warning_stddev", "14"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.http_latency_critical_stddev", "15"),
		resource.TestCheckResourceAttr(testDS, "settings.0.health_settings.0.http_latency_warning_stddev", "16"),
		resource.TestCheckNoResourceAttr(testDS, "settings.0.monitoring_settings.0.activation_grace_period"),
		resource.TestCheckNoResourceAttr(testDS, "settings.0.monitoring_settings.0.activation_time_unit"),
		resource.TestCheckNoResourceAttr(testDS, "settings.0.monitoring_settings.0.activation_time_window"),
		resource.TestCheckNoResourceAttr(testDS, "settings.0.monitoring_settings.0.activation_times"),
		resource.TestCheckResourceAttr(testDS, "settings.0.monitoring_settings.0.notification_channels.#", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.monitoring_settings.0.notification_channels.0", "nc-one"),
		resource.TestCheckResourceAttr(testDS, "settings.0.monitoring_settings.0.notification_channels.1", "nc-two"),
		resource.TestCheckResourceAttr(testDS, "settings.0.ping.0.period", "60"),
		resource.TestCheckResourceAttr(testDS, "settings.0.ping.0.count", "5"),
		resource.TestCheckResourceAttr(testDS, "settings.0.ping.0.expiry", "3000"),
		resource.TestCheckResourceAttr(testDS, "settings.0.ping.0.delay", "4"),
		resource.TestCheckResourceAttr(testDS, "settings.0.trace.0.period", "60"),
		resource.TestCheckResourceAttr(testDS, "settings.0.trace.0.count", "3"),
		resource.TestCheckResourceAttr(testDS, "settings.0.trace.0.protocol", "udp"),
		resource.TestCheckResourceAttr(testDS, "settings.0.trace.0.port", "33434"),
		resource.TestCheckResourceAttr(testDS, "settings.0.trace.0.expiry", "22500"),
		resource.TestCheckResourceAttr(testDS, "settings.0.trace.0.limit", "30"),
		resource.TestCheckResourceAttr(testDS, "settings.0.trace.0.delay", "7"),
		resource.TestCheckResourceAttr(testDS, "settings.0.port", "443"),
		resource.TestCheckResourceAttr(testDS, "settings.0.protocol", "icmp"),
		resource.TestCheckResourceAttr(testDS, "settings.0.family", "IP_FAMILY_DUAL"),
		resource.TestCheckResourceAttr(testDS, "settings.0.servers.#", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.servers.0", "192.0.2.1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.servers.1", "192.0.2.2"),
		resource.TestCheckNoResourceAttr(testDS, "settings.0.use_local_ip"),
		resource.TestCheckNoResourceAttr(testDS, "settings.0.reciprocal"),
		resource.TestCheckResourceAttr(testDS, "settings.0.rollup_level", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.http.0.period", "1"),
		resource.TestCheckResourceAttr(testDS, "settings.0.http.0.expiry", "2"),
		resource.TestCheckResourceAttr(testDS, "settings.0.http.0.method", "GET"),
		resource.TestCheckResourceAttr(testDS, "settings.0.http.0.headers.dummy-header-key", "dummy-header-value"),
		resource.TestCheckResourceAttr(testDS, "settings.0.http.0.body", "dummy-body"),
		resource.TestCheckResourceAttr(testDS, "settings.0.http.0.ignore_tls_errors", "true"),
		resource.TestCheckResourceAttr(testDS, "settings.0.http.0.css_selectors.dummy-selector-key", "dummy-selector-value"),
		resource.TestCheckResourceAttr(testDS, "expires_on", "2021-04-08T12:24:19.765Z"),
		resource.TestCheckResourceAttr(testDS, "cdate", "2021-04-08T12:24:18.75Z"),
		resource.TestCheckResourceAttr(testDS, "edate", "2021-04-13T09:20:23.819Z"),
		resource.TestCheckResourceAttr(testDS, "created_by.0.id", "dummy-id"),
		resource.TestCheckResourceAttr(testDS, "created_by.0.email", "dummy-email"),
		resource.TestCheckResourceAttr(testDS, "created_by.0.full_name", "dummy-fn"),
		resource.TestCheckResourceAttr(testDS, "last_updated_by.0.id", "dummy-id-2"),
		resource.TestCheckResourceAttr(testDS, "last_updated_by.0.email", "dummy-email-2"),
		resource.TestCheckResourceAttr(testDS, "last_updated_by.0.full_name", "dummy-fn-2"),
	)
}

const (
	testDS               = "data.kentik-synthetics_test.github-test"
	dataSourceTestConfig = `
		provider "kentik-synthetics" {
			log_payloads = true
		}

		data "kentik-synthetics_test" "github-test" {
			id = "1"
		}
	`
)
