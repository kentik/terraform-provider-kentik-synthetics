// +build tools

package tools

import (
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs" // documentation generation
	_ "github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver" // used for tests
	_ "mvdan.cc/gofumpt"
)
