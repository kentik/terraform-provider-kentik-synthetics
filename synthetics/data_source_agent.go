package synthetics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

func dataSourceAgent() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing single synthetic test agent",
		ReadContext: dataSourceAgentRead,
		Schema:      makeAgentSchema(readSingle),
	}
}

func dataSourceAgentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	agentID := d.Get("id").(string)
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceApi.AgentGet(ctx, agentID).Execute()
	if err != nil {
		return detailedDiagError("failed to read agent", err, httpResp)
	}

	d.SetId(resp.Agent.GetId())

	for k, v := range agentToMap(resp.Agent) {
		if err = d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
