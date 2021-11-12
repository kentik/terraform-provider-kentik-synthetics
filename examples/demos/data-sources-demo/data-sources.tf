data "kentik-synthetics_agents" "agents" {}

output "agents_list" {
  value = data.kentik-synthetics_agents.agents
}

data "kentik-synthetics_agent" "agent" {
  id = "1207" # ID of existing synthetics agent
}

output "agent_get" {
  value = data.kentik-synthetics_agent.agent
}

data "kentik-synthetics_tests" "tests" {}

output "tests_list" {
  value = data.kentik-synthetics_tests.tests
}

data "kentik-synthetics_test" "test" {
  id = "3271" # ID of existing synthetics test
}

output "test_get" {
  value = data.kentik-synthetics_test.test
}
