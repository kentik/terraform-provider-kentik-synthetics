data "kentik-synthetics_agent" "agent" {
  id = "968" # ID of Amsterdam Agent in test-data.json
}

output "agent_get" {
  value = data.kentik-synthetics_agent.agent
}
