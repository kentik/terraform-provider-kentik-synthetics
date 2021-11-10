package synthetics

import (
	"context"
	"math"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	geo "github.com/kellydunn/golang-geo"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

const (
	itemsKey              = "items"
	invalidAgentsCountKey = "invalid_agents_count"
	latitudeKey           = "latitude"
	longitudeKey          = "longitude"
	minDistanceKey        = "min_distance"
	maxDistanceKey        = "max_distance"
)

func dataSourceAgents() *schema.Resource {
	return &schema.Resource{
		Description: "Data source representing list of synthetic test agents",
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
			minDistanceKey: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			maxDistanceKey: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func dataSourceAgentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.AgentsList(ctx).Execute()
	if err != nil {
		return detailedDiagError("failed to read agents", err, httpResp)
	}

	if resp.Agents != nil {
		agents := filterAgents(*resp.Agents, d)
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
	d.SetId(strconv.Itoa(int(time.Now().Unix())))
	return nil
}

func agentsToMaps(agents []synthetics.V202101beta1Agent) []interface{} {
	result := make([]interface{}, len(agents))
	for i := range agents {
		result[i] = agentToMap(&agents[i])
	}
	return result
}

func filterAgents(agents []synthetics.V202101beta1Agent, d *schema.ResourceData) []synthetics.V202101beta1Agent {
	lat, latExists := d.GetOk(latitudeKey)
	lon, lonExists := d.GetOk(longitudeKey)
	minDist, minDistExists := d.GetOk(minDistanceKey)
	maxDist, maxDistExists := d.GetOk(maxDistanceKey)

	if !minDistExists {
		minDist = 0.0
	}
	if !maxDistExists {
		maxDist = math.Inf(0)
	}

	if latExists && lonExists {
		agents = filterAgentsByDistance(agents, lat.(float64), lon.(float64), minDist.(float64), maxDist.(float64))
	}

	return agents
}

func filterAgentsByDistance(agents []synthetics.V202101beta1Agent,
	lat float64, long float64, minDist float64, maxDist float64) []synthetics.V202101beta1Agent {
	var filteredAgents []synthetics.V202101beta1Agent
	referencePoint := geo.NewPoint(lat, long)
	for _, agent := range agents {
		agentCoordinates := geo.NewPoint(*agent.Lat, *agent.Long)
		dist := referencePoint.GreatCircleDistance(agentCoordinates)
		if dist >= minDist && dist <= maxDist {
			filteredAgents = append(filteredAgents, agent)
		}
	}

	return filteredAgents
}
