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
	itemsKey              = "items"
	invalidAgentsCountKey = "invalid_agents_count"
)

func dataSourceAgents() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing list of synthetic test agents",
		ReadContext: dataSourceAgentsRead,
		Schema: map[string]*schema.Schema{
			itemsKey: {
				Type:     schema.TypeList,
				Computed: computedOnRead(readList),
				Elem: &schema.Resource{
					Schema: makeAgentSchema(readList),
				},
			},
			invalidAgentsCountKey: {
				Type:     schema.TypeInt,
				Computed: computedOnRead(readList),
			},
		},
	}
}

func dataSourceAgentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// use UNIX time as ID to force list update every time Terraform asks for the list
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceApi.AgentsList(ctx).Execute()
	if err != nil {
		return detailedDiagError("failed to read agents", err, httpResp)
	}

	if resp.Agents != nil {
		err = d.Set(itemsKey, agentsToMaps(*resp.Agents))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if resp.InvalidAgentsCount != nil {
		err = d.Set(invalidAgentsCountKey, resp.InvalidAgentsCount)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func agentsToMaps(agents []synthetics.V202101beta1Agent) []interface{} {
	result := make([]interface{}, len(agents))
	for i, e := range agents {
		result[i] = agentToMap(&e)
	}
	return result
}
