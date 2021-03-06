package synthetics

// Note the file is named "data_source_test_impl.go" instead of "data_source_test.go", so that
// Go toolchain does not treat it as test file (because of _test.go suffix)

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/kentikapi"
)

func dataSourceTest() *schema.Resource {
	return &schema.Resource{
		Description: "Data source representing single synthetic test",
		ReadContext: dataSourceTestRead,
		Schema:      makeTestSchema(readSingle),
	}
}

func dataSourceTestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Get Test Kentik API request", map[string]interface{}{"ID": d.Get(idKey).(string)})
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
		TestGet(ctx, d.Get(idKey).(string)).
		Execute()
	tflog.Debug(ctx, "Get Test Kentik API response", map[string]interface{}{"response": resp})
	if err != nil {
		return detailedDiagError("Failed to read test", err, httpResp)
	}

	for k, v := range testToMap(resp.Test) {
		if err = d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(resp.Test.GetId())
	return nil
}
