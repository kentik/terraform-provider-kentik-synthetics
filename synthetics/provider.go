package synthetics

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

const (
	apiURLKey      = "api_url"
	emailKey       = "email"
	tokenKey       = "token"
	logPayloadsKey = "log_payloads"
)

// NewProvider returns new Synthetics provider.
func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			apiURLKey: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_URL", nil),
				Description: "Synthetics API server URL (optional). Can also be specified with KTAPI_URL environment variable.",
			},
			emailKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_EMAIL", nil),
				Description: "Authorization email (required). Can also be specified with KTAPI_AUTH_EMAIL environment variable.",
			},
			tokenKey: {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_TOKEN", nil),
				Description: "Authorization token (required). Can also be specified with KTAPI_AUTH_TOKEN environment variable.",
			},
			logPayloadsKey: {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_LOG_PAYLOADS", false),
				Description: "Log payloads flag enables verbose debug logs of requests and responses (optional). " +
					"Can also be specified with KTAPI_LOG_PAYLOADS environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kentik-synthetics_test": resourceTest(),
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
	cfg := kentikapi.Config{
		SyntheticsAPIURL: getURL(d),
		AuthEmail:        d.Get(emailKey).(string),
		AuthToken:        d.Get(tokenKey).(string),
		LogPayloads:      d.Get(logPayloadsKey).(bool),
	}
	log.Printf("[DEBUG] Creating Kentik API client with config: %+v", stripSensitiveData(cfg))

	return kentikapi.NewClient(cfg), nil
}

func getURL(d *schema.ResourceData) string {
	var url string
	apiURL, ok := d.GetOk(apiURLKey)
	if ok {
		url = apiURL.(string) //nolint: errcheck, forcetypeassert // type enforced by TF schema
	}
	return url
}

func stripSensitiveData(cfg kentikapi.Config) kentikapi.Config {
	cfg.AuthToken = "<stripped>"
	return cfg
}
