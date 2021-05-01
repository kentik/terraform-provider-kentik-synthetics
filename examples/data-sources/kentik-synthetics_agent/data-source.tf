data "kentik-synthetics_agent" "agent" {
  id = "824"
}

output "agent_get" {
  value = data.kentik-synthetics_agent.agent
}
