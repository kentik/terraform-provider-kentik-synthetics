package synthetics

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

const (
	invalidTestsCountKey = "invalid_tests_count"
)

func dataSourceTests() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing list of synthetic test tests",
		ReadContext: dataSourceTestsRead,
		Schema: map[string]*schema.Schema{
			itemsKey: {
				Type:     schema.TypeList,
				Computed: computedOnRead(readList),
				Elem: &schema.Resource{
					Schema: makeTestSchema(readList),
				},
			},
			invalidTestsCountKey: {
				Type:     schema.TypeInt,
				Computed: computedOnRead(readList),
			},
		},
	}
}

func dataSourceTestsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// use UNIX time as ID to force list update every time Terraform asks for the list
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceApi.TestsList(ctx).Execute()
	if err != nil {
		return detailedDiagError("failed to read tests", err, httpResp)
	}

	if resp.Tests != nil {
		err = d.Set(itemsKey, testsToMaps(*resp.Tests))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if resp.InvalidTestsCount != nil {
		err = d.Set(invalidTestsCountKey, resp.InvalidTestsCount)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func testsToMaps(tests []synthetics.V202101beta1Test) []interface{} {
	result := make([]interface{}, len(tests))
	for i, e := range tests {
		result[i] = testToMap(&e)
	}
	return result
}
