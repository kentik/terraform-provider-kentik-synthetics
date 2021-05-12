data "kentik-synthetics_tests" "tests" {}

output "tests_list" {
  value = data.kentik-synthetics_tests.tests
}
