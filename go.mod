module github.com/kentik/terraform-provider-kentik-synthetics

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/kentik/community_sdk_golang/apiv6 v0.0.0-20210503202321-1dab307d87ff
)

replace github.com/kentik/community_sdk_golang/apiv6 => github.com/Fenthick/community_sdk_golang/apiv6 v0.0.0-20210531094034-d65e4592bcb6
