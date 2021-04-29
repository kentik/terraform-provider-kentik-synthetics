module github.com/kentik/terraform-provider-kentik-synthetics

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/kentik/community_sdk_golang/apiv6 v0.0.0-00010101000000-000000000000
)

replace github.com/kentik/community_sdk_golang/apiv6 => github.com/mateuszmidor/community_sdk_golang/apiv6 v0.0.0-20210422065113-ea13968f2c8a
