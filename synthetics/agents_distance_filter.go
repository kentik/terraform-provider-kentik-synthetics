package synthetics

import (
	geo "github.com/kellydunn/golang-geo"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

func filterAgentsByDistance(agents []synthetics.V202101beta1Agent,
	lat float64, long float64, distConstr float64) []synthetics.V202101beta1Agent {
	var filteredAgents []synthetics.V202101beta1Agent
	p := geo.NewPoint(lat, long)
	for _, agent := range agents {
		p2 := geo.NewPoint(*agent.Lat, *agent.Long)
		dist := p.GreatCircleDistance(p2)
		if dist < distConstr {
			filteredAgents = append(filteredAgents, agent)
		}
	}

	return filteredAgents
}
