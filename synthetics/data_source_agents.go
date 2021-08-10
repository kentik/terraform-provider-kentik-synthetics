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
	latitudeKey           = "latitude"
	longitudeKey          = "longitude"
	distanceKey           = "distance"
	nameSubstringKey      = "name_substring"
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
			latitudeKey: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			longitudeKey: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			distanceKey: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			nameSubstringKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceAgentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceApi.AgentsList(ctx).Execute()
	if err != nil {
		return detailedDiagError("failed to read agents", err, httpResp)
	}

	if resp.Agents != nil {
		lat, latExists := d.GetOk(latitudeKey)
		lon, lonExists := d.GetOk(longitudeKey)
		dist, distExists := d.GetOk(distanceKey)
		agents := *resp.Agents

		if latExists && lonExists && distExists {
			agents = filterAgentsByDistance(*resp.Agents, lat.(float64), lon.(float64), dist.(float64))
		}

		nameSubstring, nameSubstringExists := d.GetOk(nameSubstringKey)
		if nameSubstringExists {
			agents = filterAgentsByName(agents, nameSubstring.(string))
		}

		agentsMap := agentsToMaps(agents)
		err = d.Set(itemsKey, agentsMap)
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

	// use UNIX time as ID to force list update every time Terraform asks for the list
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func agentsToMaps(agents []synthetics.V202101beta1Agent) []interface{} {
	result := make([]interface{}, len(agents))
	for i := range agents {
		result[i] = agentToMap(&agents[i])
	}
	return result
}
