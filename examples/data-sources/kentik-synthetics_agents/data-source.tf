data "kentik-synthetics_agents" "agents" {}

output "agents_list" {
  value = data.kentik-synthetics_agents.agents
}
