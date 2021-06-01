package synthetics

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

func dataSourceAgent() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing single synthetic test agent",
		ReadContext: dataSourceAgentRead,
		Schema:      makeAgentSchema(readSingle),
	}
}

func dataSourceAgentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceApi.
		AgentGet(ctx, d.Get(idKey).(string)).
		Execute()
	if err != nil {
		return detailedDiagError("failed to read agent", err, httpResp)
	}

	for k, v := range agentToMap(resp.Agent) {
		if err = d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(resp.Agent.GetId())
	return nil
}

func agentToMap(obj *synthetics.V202101beta1Agent) map[string]interface{} {
	if obj == nil {
		return nil
	}

	m := make(map[string]interface{})
	m[idKey] = obj.Id
	m["name"] = obj.Name
	m["status"] = obj.Status
	m["alias"] = obj.Alias
	m["type"] = obj.Type
	m["os"] = obj.Os
	m["ip"] = obj.Ip
	m["lat"] = obj.Lat
	m["long"] = obj.Long
	if obj.LastAuthed != nil {
		m["last_authed"] = obj.LastAuthed.Format(time.RFC3339Nano)
	}
	m["family"] = obj.Family
	m["asn"] = obj.Asn
	m["site_id"] = obj.SiteId
	m["version"] = obj.Version
	m["challenge"] = obj.Challenge
	m["city"] = obj.City
	m["region"] = obj.Region
	m["country"] = obj.Country
	m["test_ids"] = obj.TestIds
	m["local_ip"] = obj.LocalIp
	return m
}
