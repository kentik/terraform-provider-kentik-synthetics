package synthetics

import (
	geo "github.com/kellydunn/golang-geo"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
	"strings"
)

func filterAgentsByDistance(agents []synthetics.V202101beta1Agent,
	lat float64, long float64, distConstr float64) []synthetics.V202101beta1Agent {
	var filteredAgents []synthetics.V202101beta1Agent
	referencePoint := geo.NewPoint(lat, long)
	for _, agent := range agents {
		agentCoordinates := geo.NewPoint(*agent.Lat, *agent.Long)
		dist := referencePoint.GreatCircleDistance(agentCoordinates)
		if dist <= distConstr {
			filteredAgents = append(filteredAgents, agent)
		}
	}

	return filteredAgents
}

func filterAgentsByName(agents []synthetics.V202101beta1Agent, nameSubstring string) []synthetics.V202101beta1Agent {
	var filteredAgents []synthetics.V202101beta1Agent
	for _, agent := range agents {
		if strings.Contains(*agent.Name, nameSubstring) {
			filteredAgents = append(filteredAgents, agent)
		}
	}

	return filteredAgents
}
