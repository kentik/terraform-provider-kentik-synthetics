package synthetics_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Note: values checked in tests below are provided by stub API Server from test-data.json (running in background)

func TestDataSourceAgent(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config: dataSourceAgentConfig,
			Check:  checkDataSourceAgent(),
		}},
	})
}

func checkDataSourceAgent() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(agentDS, "id", "968"),
		resource.TestCheckResourceAttr(agentDS, "name", "global-agent"),
		resource.TestCheckResourceAttr(agentDS, "status", "AGENT_STATUS_WAIT"),
		resource.TestCheckResourceAttr(agentDS, "alias", "probe-4-ams-1"),
		resource.TestCheckResourceAttr(agentDS, "type", "global"),
		resource.TestCheckResourceAttr(agentDS, "os",
			"Linux probe-4-ams-1 4.9.0-12-amd64 #1 SMP Debian 4.9.210-1+deb9u1 (2020-06-07) x86_64",
		),
		resource.TestCheckResourceAttr(agentDS, "ip", "95.179.136.58"),
		resource.TestCheckResourceAttr(agentDS, "lat", "52.374031"),
		resource.TestCheckResourceAttr(agentDS, "long", "4.88969"),
		resource.TestCheckResourceAttr(agentDS, "last_authed", "2020-07-09T21:29:43.826Z"),
		resource.TestCheckResourceAttr(agentDS, "family", "IP_FAMILY_DUAL"),
		resource.TestCheckResourceAttr(agentDS, "asn", "20473"),
		resource.TestCheckResourceAttr(agentDS, "site_id", "0"),
		resource.TestCheckResourceAttr(agentDS, "version", "0.0.2"),
		resource.TestCheckResourceAttr(agentDS, "challenge", ""),
		resource.TestCheckResourceAttr(agentDS, "city", "Amsterdam"),
		resource.TestCheckResourceAttr(agentDS, "region", "Noord-Holland"),
		resource.TestCheckResourceAttr(agentDS, "country", "Netherlands"),
		resource.TestCheckResourceAttr(agentDS, "test_ids.#", "3"),
		resource.TestCheckResourceAttr(agentDS, "test_ids.0", "13"),
		resource.TestCheckResourceAttr(agentDS, "test_ids.1", "133"),
		resource.TestCheckResourceAttr(agentDS, "test_ids.2", "1337"),
		resource.TestCheckResourceAttr(agentDS, "local_ip", ""),
	)
}

const (
	agentDS               = "data.kentik-synthetics_agent.amsterdam-agent"
	dataSourceAgentConfig = `
		data "kentik-synthetics_agent" "amsterdam-agent" {
			id = "968"
		}
	`
)
