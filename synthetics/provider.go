package synthetics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

const (
	apiURLKey = "api_url"
	emailKey  = "email"
	tokenKey  = "token"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

// NewProvider returns new Synthetics provider.
func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			apiURLKey: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_URL", nil),
				Description: "Custom API server url can be specified either by api_url attribute or KTAPI_URL environment variable. If not specified, default of <https://synthetics.api.kentik.com> will be used",
			},
			emailKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_EMAIL", nil),
				Description: "Authorization. Either email attribute or KTAPI_AUTH_EMAIL environment variable is required",
			},
			tokenKey: {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_TOKEN", nil),
				Description: "Authorization. Either token attribute or KTAPI_AUTH_TOKEN environment variable is required",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"kentik-synthetics_agent":  dataSourceAgent(),
			"kentik-synthetics_agents": dataSourceAgents(),
			"kentik-synthetics_tests":  dataSourceTests(),
			"kentik-synthetics_test":   dataSourceTest(),
		},
		ConfigureContextFunc: configure,
	}
}

func configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	email := d.Get(emailKey).(string)
	token := d.Get(tokenKey).(string)

	apiURL, ok := d.GetOk(apiURLKey)
	if !ok {
		return newClient(email, token, ""), nil
	}

	return newClient(email, token, apiURL.(string)), nil
}

func newClient(email, token, url string) *kentikapi.Client {
	return kentikapi.NewClient(kentikapi.Config{
		SyntheticsAPIURL: url,
		AuthEmail:        email,
		AuthToken:        token,
	})
}
