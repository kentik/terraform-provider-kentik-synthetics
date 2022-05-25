package synthetics

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

// TODO(dfurman): provide descriptions, when they are specified in the OpenAPI definitions

// makeTestSchema omits following internal attributes: device_id.
func makeTestSchema(mode schemaMode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Required: requiredOnReadSingle(mode),
			Computed: computedOnCreateAndReadList(mode),
		},
		"name": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{
					"ip", "hostname", "agent", "network_grid", "application_mesh", "flow", "url", "page_load", "dns", "dns_grid",
				},
				false))),
		},
		"status": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{
					string(synthetics.V202101BETA1TESTSTATUS_ACTIVE),
					string(synthetics.V202101BETA1TESTSTATUS_PAUSED),
					string(synthetics.V202101BETA1TESTSTATUS_DELETED),
				}, false))),
		},
		"settings": makeTestSettingsSchema(mode),
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"cdate": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"edate": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_by":      makeUserInfoSchema(),
		"last_updated_by": makeUserInfoSchema(),
	}
}

// makeTestSettingsSchema omits following internal attributes: use_local_ip, reciprocal.
func makeTestSettingsSchema(mode schemaMode) *schema.Schema {
	return makeRequiredNestedObjectSchema(mode, map[string]*schema.Schema{
		"hostname":     makeTestHostnameSchema(mode),
		"ip":           makeTestIPSchema(mode),
		"agent":        makeTestAgentSchema(mode),
		"flow":         makeTestFlowSchema(mode),
		"site":         makeTestSiteSchema(mode),
		"tag":          makeTestTagSchema(mode),
		"dns":          makeTestDNSSchema(mode),
		"url":          makeTestURLSchema(mode),
		"network_grid": makeTestNetworkGridSchema(mode),
		"page_load":    makeTestPageLoadSchema(mode),
		"dns_grid":     makeTestDNSGridSchema(mode),
		"agent_ids": {
			Type:     schema.TypeList,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"period": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntBetween(1, 900),
		},
		"count": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntAtLeast(0),
		},
		"expiry": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntAtLeast(0),
		},
		"limit": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntAtLeast(0),
		},
		"tasks": {
			Type:     schema.TypeList,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"health_settings":     makeTestHealthSettingsSchema(mode),
		"monitoring_settings": makeTestMonitoringSettingsSchema(mode),
		"ping":                makeTestPingSchema(mode),
		"trace":               makeTestTraceSchema(mode),
		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IsPortNumberOrZero,
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{"icmp", "tcp", "udp"}, false))),
		},
		"family": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{
					string(synthetics.V202101BETA1IPFAMILY_UNSPECIFIED),
					string(synthetics.V202101BETA1IPFAMILY_V4),
					string(synthetics.V202101BETA1IPFAMILY_V6),
					string(synthetics.V202101BETA1IPFAMILY_DUAL),
				}, false))),
		},
		"servers": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"rollup_level": {
			Type:     schema.TypeInt,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"http": makeTestHTTPSchema(mode),
	})
}

func makeTestHostnameSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestIPSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"targets": {
			Type:     schema.TypeList,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
		},
	})
}

func makeTestAgentSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestFlowSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"target_refresh_interval_millis": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntBetween(3600000, 604800000),
		},
		"max_tasks": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{"asn", "cdn", "country", "region", "city"}, false))),
		},
		"inet_direction": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{"src", "dst"}, false))),
		},
		"direction": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{"src", "dst"}, false))),
		},
	})
}

func makeTestSiteSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestTagSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestDNSSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{
					string(synthetics.V202101BETA1DNSRECORD_UNSPECIFIED),
					string(synthetics.V202101BETA1DNSRECORD_A),
					string(synthetics.V202101BETA1DNSRECORD_AAAA),
					string(synthetics.V202101BETA1DNSRECORD_CNAME),
					string(synthetics.V202101BETA1DNSRECORD_DNAME),
					string(synthetics.V202101BETA1DNSRECORD_NS),
					string(synthetics.V202101BETA1DNSRECORD_MX),
					string(synthetics.V202101BETA1DNSRECORD_PTR),
					string(synthetics.V202101BETA1DNSRECORD_SOA),
				}, false))),
		},
	})
}

func makeTestURLSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:             schema.TypeString,
			Required:         requiredOnCreate(mode),
			Computed:         computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.IsURLWithHTTPorHTTPS)),
		},
	})
}

func makeTestNetworkGridSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"targets": {
			Type:     schema.TypeList,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
		},
	})
}

func makeTestPageLoadSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:             schema.TypeString,
			Required:         requiredOnCreate(mode),
			Computed:         computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.IsURLWithHTTPorHTTPS)),
		},
	})
}

func makeTestDNSGridSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"targets": {
			Type:     schema.TypeList,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{
					string(synthetics.V202101BETA1DNSRECORD_UNSPECIFIED),
					string(synthetics.V202101BETA1DNSRECORD_A),
					string(synthetics.V202101BETA1DNSRECORD_AAAA),
					string(synthetics.V202101BETA1DNSRECORD_CNAME),
					string(synthetics.V202101BETA1DNSRECORD_DNAME),
					string(synthetics.V202101BETA1DNSRECORD_NS),
					string(synthetics.V202101BETA1DNSRECORD_MX),
					string(synthetics.V202101BETA1DNSRECORD_PTR),
					string(synthetics.V202101BETA1DNSRECORD_SOA),
				}, false))),
		},
	})
}

func makeTestHealthSettingsSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"latency_critical": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"latency_warning": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"packet_loss_critical": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100),
		},
		"packet_loss_warning": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100),
		},
		"jitter_critical": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"jitter_warning": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"http_latency_critical": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100),
		},
		"http_latency_warning": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100),
		},
		"http_valid_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntAtLeast(0),
			},
		},
		"dns_valid_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntAtLeast(0),
			},
		},
		"latency_critical_stddev": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100),
		},
		"latency_warning_stddev": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100),
		},
		"jitter_critical_stddev": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"jitter_warning_stddev": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"http_latency_critical_stddev": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"http_latency_warning_stddev": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
	})
}

// makeTestMonitoringSettingsSchema omits following internal attributes: activation_grace_period, activation_time_unit,
// activation_time_window, activation_times.
func makeTestMonitoringSettingsSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"notification_channels": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	})
}

func makeTestPingSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"period": {
			Type:     schema.TypeFloat,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(1, 10),
		},
		"expiry": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100000),
		},
		"delay": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatAtLeast(0),
		},
	})
}

func makeTestTraceSchema(mode schemaMode) *schema.Schema {
	return makeRequiredNestedObjectSchema(mode, map[string]*schema.Schema{
		"period": {
			Type:     schema.TypeFloat,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 10),
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{"icmp", "tcp", "udp"}, false))),
		},
		"port": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 100000),
		},
		"limit": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatBetween(0, 255),
		},
		"delay": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.FloatAtLeast(0),
		},
	})
}

func makeTestHTTPSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"period": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntAtLeast(0),
		},
		"expiry": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntBetween(5000, 60000),
		},
		"method": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{"GET", "PATCH", "POST", "PUT"}, false))),
		},
		"headers": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"body": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"ignore_tls_errors": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"css_selectors": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: computedOnRead(mode),
		},
	})
}

func makeUserInfoSchema() *schema.Schema {
	return makeReadOnlyNestedObjectSchema(map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		"email": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"full_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	})
}
