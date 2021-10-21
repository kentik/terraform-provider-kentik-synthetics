package synthetics

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

const (
	apiURLKey      = "api_url"
	emailKey       = "email"
	tokenKey       = "token"
	retryKey       = "retry"
	maxAttemptsKey = "max_attempts"
	minDelayKey    = "min_delay"
	maxDelayKey    = "max_delay"
	logPayloadsKey = "log_payloads"

	defaultMaxAttempts = 100
	defaultMinDelay    = "1s"
	defaultMaxDelay    = "5m"
)

// NewProvider returns new Synthetics provider.
func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: providerSchema(),
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

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		retryKey: {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			ForceNew:    true,
			Description: "Configuration for API client retry mechanism",
			Elem: &schema.Resource{
				Schema: retryProperties(),
			},
		},
		logPayloadsKey: {
			Type:        schema.TypeBool,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KTAPI_LOG_PAYLOADS", false),
			Description: "Log payloads flag enables verbose debug logs of requests and responses (optional). " +
				"Can also be specified with KTAPI_LOG_PAYLOADS environment variable.",
		},
	}
}

// retryProperties groups retry properties.
// Note that, default values of schema.EnvDefaultFunc are not used, because they are not applied when user does not pass
// retry block at all.
func retryProperties() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		maxAttemptsKey: {
			Type:        schema.TypeInt,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KTAPI_RETRY_MAX_ATTEMPTS", nil),
			Description: "Maximum number of request retry attempts. " +
				"Minimum valid value: 1 (0 fallbacks to default). Default: 100. " +
				"Can also be specified with KTAPI_RETRY_MAX_ATTEMPTS environment variable.",
		},
		minDelayKey: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KTAPI_RETRY_MIN_DELAY", nil),
			Description: "Minimum delay before request retry. " +
				"Expected Go time duration format, e.g. 1s (see: <https://pkg.go.dev/time#ParseDuration>). " +
				"Default: 1s (1 second). " +
				"Can also be specified with KTAPI_RETRY_MIN_DELAY environment variable.",
		},
		maxDelayKey: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KTAPI_RETRY_MAX_DELAY", nil),
			Description: "Maximum delay before request retry. " +
				"Expected Go time duration format, e.g. 1s (see: <https://pkg.go.dev/time#ParseDuration>). " +
				"Default: 5m (5 minutes). " +
				"Can also be specified with KTAPI_RETRY_MAX_DELAY environment variable.",
		},
	}
}

func configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	rc, err := getRetryConfig(d)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	cfg := kentikapi.Config{
		SyntheticsAPIURL: getAPIURL(d),
		AuthEmail:        d.Get(emailKey).(string),
		AuthToken:        d.Get(tokenKey).(string),
		RetryCfg:         rc,
		LogPayloads:      d.Get(logPayloadsKey).(bool),
	}

	strippedCfg := stripSensitiveData(cfg)
	cfgJSON, _ := json.Marshal(strippedCfg) //nolint: errcheck
	log.Printf("Creating Kentik API client with config: %+v, JSON: %v", strippedCfg, string(cfgJSON))

	return kentikapi.NewClient(cfg), nil
}

func getRetryConfig(d *schema.ResourceData) (kentikapi.RetryConfig, error) {
	log.Printf("Getting retry config: %v", d.Get(retryKey))
	retryCfg, err := getObjectFromNestedResourceData(d.Get(retryKey))
	if err != nil {
		return kentikapi.RetryConfig{}, fmt.Errorf("get retry configuration: %v", err)
	}

	maxAttempts, ok := retryCfg[maxAttemptsKey].(int)
	if !ok || maxAttempts == 0 {
		maxAttempts = defaultMaxAttempts
	}

	rawMinDelay, ok := retryCfg[minDelayKey].(string)
	if !ok || rawMinDelay == "" {
		rawMinDelay = defaultMinDelay
	}
	minDelay, err := time.ParseDuration(rawMinDelay)
	if err != nil {
		return kentikapi.RetryConfig{}, fmt.Errorf("parse %v duration: %v", minDelayKey, err)
	}

	rawMaxDelay, ok := retryCfg[maxDelayKey].(string)
	if !ok || rawMaxDelay == "" {
		rawMaxDelay = defaultMaxDelay
	}
	maxDelay, err := time.ParseDuration(rawMaxDelay)
	if err != nil {
		return kentikapi.RetryConfig{}, fmt.Errorf("parse %v duration: %v", maxDelayKey, err)
	}

	return kentikapi.RetryConfig{
		MaxAttempts: &maxAttempts,
		MinDelay:    &minDelay,
		MaxDelay:    &maxDelay,
	}, nil
}

func getAPIURL(d *schema.ResourceData) string {
	apiURL, ok := d.GetOk(apiURLKey)
	if ok {
		return apiURL.(string)
	}
	return ""
}

func stripSensitiveData(cfg kentikapi.Config) kentikapi.Config {
	cfg.AuthToken = "stripped"
	return cfg
}
