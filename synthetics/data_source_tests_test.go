package synthetics_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Note: values checked in tests below are provided by stub API Server from test-data.json (running in background)

func TestDataSourceTests(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config: dataSourceTestsConfig,
			Check:  checkDataSourceTests(),
		}},
	})
}

func checkDataSourceTests() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		// nested properties, such as settings, are verified in TestDataSourceTest
		resource.TestCheckResourceAttr(testsDS, "items.0.id", "1"),
		resource.TestCheckResourceAttr(testsDS, "items.0.name", "github.com ip test"),
		resource.TestCheckResourceAttr(testsDS, "items.0.type", "ip"),
		resource.TestCheckResourceAttr(testsDS, "items.0.device_id", "75702"),
		resource.TestCheckResourceAttr(testsDS, "items.0.status", "TEST_STATUS_PAUSED"),
		resource.TestCheckResourceAttrSet(testsDS, "items.0.settings.0.count"),
		resource.TestCheckResourceAttr(testsDS, "items.0.expires_on", "2021-04-08T12:24:19.765Z"),
		resource.TestCheckResourceAttr(testsDS, "items.0.cdate", "2021-04-08T12:24:18.75Z"),
		resource.TestCheckResourceAttr(testsDS, "items.0.edate", "2021-04-13T09:20:23.819Z"),
		resource.TestCheckResourceAttrSet(testsDS, "items.0.created_by.0.id"),
		resource.TestCheckResourceAttrSet(testsDS, "items.0.last_updated_by.0.id"),

		resource.TestCheckResourceAttr(testsDS, "items.1.id", "2"),
		resource.TestCheckResourceAttr(testsDS, "items.1.name", "metoffice.gov.uk ip test"),
		resource.TestCheckResourceAttr(testsDS, "items.1.type", "ip"),
		resource.TestCheckResourceAttr(testsDS, "items.1.device_id", "75702"),
		resource.TestCheckResourceAttr(testsDS, "items.1.status", "TEST_STATUS_ACTIVE"),
		resource.TestCheckResourceAttrSet(testsDS, "items.1.settings.0.count"),
		resource.TestCheckResourceAttr(testsDS, "items.1.expires_on", "2021-04-08T12:00:03.097Z"),
		resource.TestCheckResourceAttr(testsDS, "items.1.cdate", "2021-04-08T12:00:02.084Z"),
		resource.TestCheckResourceAttr(testsDS, "items.1.edate", "2021-04-16T08:14:35.217Z"),
		resource.TestCheckResourceAttr(testsDS, "items.1.created_by.0.id", ""),
		resource.TestCheckResourceAttr(testsDS, "items.1.last_updated_by.0.id", ""),

		// TODO(dfurman): modify stub API server to return invalid_tests_count key
		// resource.TestCheckResourceAttr(testsDS, "invalid_tests_count", "5"),
	)
}

const (
	testsDS               = "data.kentik-synthetics_tests.dummy-tests"
	dataSourceTestsConfig = `
		provider "kentik-synthetics" {}
		data "kentik-synthetics_tests" "dummy-tests" {}
	`
)
