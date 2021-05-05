data "kentik-synthetics_test" "test" {
  id = "1" # ID of "github.com ip test" in test-data.json
}

output "test_get" {
  value = data.kentik-synthetics_test.test
}
