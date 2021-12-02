//go:build tools

package tools

import (
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"    // documentation generation
	_ "github.com/kentik/community_sdk_golang/kentikapi/fakeapiserver" // used for tests
)
