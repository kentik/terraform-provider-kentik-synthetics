package synthetics_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Note: values checked in tests below are provided by stub API Server from test-data.json (running in background)

func TestDataSourceAgents(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config: dataSourceAgentsConfig,
			Check:  checkDataSourceAgents(),
		}},
	})
}

func checkDataSourceAgents() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(agentsDS, "items.0.id", "968"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.name", "global-agent"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.status", "AGENT_STATUS_WAIT"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.alias", "probe-4-ams-1"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.type", "global"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.os",
			"Linux probe-4-ams-1 4.9.0-12-amd64 #1 SMP Debian 4.9.210-1+deb9u1 (2020-06-07) x86_64",
		),
		resource.TestCheckResourceAttr(agentsDS, "items.0.ip", "95.179.136.58"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.lat", "52.374031"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.long", "4.88969"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.last_authed", "2020-07-09T21:29:43.826Z"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.family", "IP_FAMILY_DUAL"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.asn", "20473"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.site_id", "0"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.version", "0.0.2"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.city", "Amsterdam"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.region", "Noord-Holland"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.country", "Netherlands"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.test_ids.#", "3"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.test_ids.0", "13"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.test_ids.1", "133"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.test_ids.2", "1337"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.local_ip", "192.168.1.2"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.cloud_vpc", "dummy-cloud-vpc"),
		resource.TestCheckResourceAttr(agentsDS, "items.0.agent_impl", "IMPLEMENT_TYPE_RUST"),

		resource.TestCheckResourceAttr(agentsDS, "items.1.id", "1717"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.name", "private-agent"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.status", "AGENT_STATUS_OK"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.alias", "ip-172-31-0-198"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.type", "private"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.os",
			"Linux ip-172-31-0-198 5.4.0-1029-aws #30-Ubuntu SMP Tue Oct 20 10:06:38 UTC 2020 x86_64",
		),
		resource.TestCheckResourceAttr(agentsDS, "items.1.ip", "18.144.28.163"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.lat", "37.774929"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.long", "-122.419418"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.last_authed", "2020-12-17T07:13:31.62Z"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.family", "IP_FAMILY_V4"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.asn", "16509"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.site_id", "7970"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.version", "0.0.9"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.city", "San Francisco"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.region", "Oregon"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.country", "US"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.test_ids.#", "0"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.local_ip", "192.168.1.2"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.cloud_vpc", "dummy-cloud-vpc"),
		resource.TestCheckResourceAttr(agentsDS, "items.1.agent_impl", "IMPLEMENT_TYPE_NODE"),

		// TODO(dfurman): modify stub API server to return invalid_agents_count key
		// resource.TestCheckResourceAttr(agentsDS, "invalid_agents_count", "5"),
	)
}

const (
	agentsDS               = "data.kentik-synthetics_agents.dummy-agents"
	dataSourceAgentsConfig = `
		provider "kentik-synthetics" {
			log_payloads = true
		}
		data "kentik-synthetics_agents" "dummy-agents" {}
	`
)
