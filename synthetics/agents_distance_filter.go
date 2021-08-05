package synthetics

import (
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
	"math"
)

func filterAgentsByDistance(agents []synthetics.V202101beta1Agent, lat float64, long float64, distConstr float64) []synthetics.V202101beta1Agent {
	var filteredAgents []synthetics.V202101beta1Agent
	for _, agent := range agents {
		dist := geoDistance(lat, long, *agent.Lat, *agent.Long, "K")
		if dist < distConstr {
			filteredAgents = append(filteredAgents, agent)
		}
	}

	return filteredAgents
}

func geoDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radLat1 := float64(PI * lat1 / 180)
	radLat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radTheta := float64(PI * theta / 180)

	dist := math.Sin(radLat1) * math.Sin(radLat2) + math.Cos(radLat1) * math.Cos(radLat2) * math.Cos(radTheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}