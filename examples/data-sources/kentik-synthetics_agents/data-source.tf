data "kentik-synthetics_agents" "agents" {
  latitude  = 50.55
  longitude = 5.5
  distance  = 1000
}

output "agents_list" {
  value = data.kentik-synthetics_agents.agents
}
