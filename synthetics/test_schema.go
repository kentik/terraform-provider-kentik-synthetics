package synthetics

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
		},
		"status": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
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
			ValidateFunc: validation.IntAtLeast(0),
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
			ValidateFunc: validation.IntAtLeast(0),
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"family": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
			ValidateDiagFunc: skipOnRead(mode, validation.ToDiagFunc(validation.StringInSlice(
				[]string{"IP_FAMILY_V4", "IP_FAMILY_V6", "IP_FAMILY_DUAL", "IP_FAMILY_UNSP", "ECIFIED"}, false))),
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
				Type: schema.TypeString,
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
			ValidateFunc: validation.IntAtLeast(0),
		},
		"max_tasks": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     computedOnRead(mode),
			ValidateFunc: validation.IntAtLeast(0),
		},
		"type": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"inet_direction": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"direction": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
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
					"DNS_RECORD_UNSPECIFIED", "DNS_RECORD_A", "DNS_RECORD_AAAA", "DNS_RECORD_CNAME",
					"DNS_RECORD_DNAME", "DNS_RECORD_NS", "DNS_RECORD_MX", "DNS_RECORD_PTR", "DNS_RECORD_SOA",
				},
				false))),
		},
	})
}

func skipOnRead(mode schemaMode, diagFunc schema.SchemaValidateDiagFunc) schema.SchemaValidateDiagFunc {
	if mode != create {
		return nil
	}
	return diagFunc
}

func makeTestURLSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
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
				Type: schema.TypeString,
			},
		},
	})
}

func makeTestPageLoadSchema(mode schemaMode) *schema.Schema {
	return makeOptionalNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
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
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
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
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"packet_loss_warning": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
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
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"http_latency_warning": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
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
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"latency_warning_stddev": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
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
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"delay": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
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
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: requiredOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"port": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"limit": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
		},
		"delay": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: computedOnRead(mode),
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
			ValidateFunc: validation.IntAtLeast(0),
		},
		"method": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: computedOnRead(mode),
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
