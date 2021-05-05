package synthetics

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func makeTestSchema(mode schemaMode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: computedOnCreateAndReadList(mode),
			Required: requiredOnReadSingle(mode),
		},
		"name": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"device_id": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"status": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
			// enumeration: "TEST_STATUS_UNSPECIFIED", TEST_STATUS_ACTIVE", "TEST_STATUS_PAUSED",
			// "TEST_STATUS_DELETED"
		},
		"settings": makeTestSettingsSchema(mode),
		"expires_on": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"cdate": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"edate": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"created_by":      makeUserInfoSchema(mode),
		"last_updated_by": makeUserInfoSchema(mode),
	}
}

func makeTestSettingsSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"hostname": makeTestHostnameSchema(mode),
		"ip":       makeTestIPSchema(mode),
		"agent":    makeTestAgentSchema(mode),
		"flow":     makeTestFlowSchema(mode),
		"site":     makeTestSiteSchema(mode),
		"tag":      makeTestTagSchema(mode),
		"dns":      makeTestDNSSchema(mode),
		"url":      makeTestURLSchema(mode),
		"agent_ids": {
			Type:     schema.TypeList,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"period": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"limit": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"tasks": {
			Type:     schema.TypeList,
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
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"protocol": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"family": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
			// enumeration: "IP_FAMILY_UNSPECIFIED", "IP_FAMILY_V4", "IP_FAMILY_V6", "IP_FAMILY_DUAL"
		},
		"servers": {
			Type:     schema.TypeList,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"target_type": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"target_value": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"use_local_ip": {
			Type:     schema.TypeBool,
			Computed: computedOnRead(mode),
		},
		"reciprocal": {
			Type:     schema.TypeBool,
			Computed: computedOnRead(mode),
		},
		"rollup_level": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestHostnameSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestIPSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"targets": {
			Type:     schema.TypeList,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	})
}

func makeTestAgentSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestFlowSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"target_refresh_interval_millis": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"max_tasks": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestSiteSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestTagSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestDNSSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestURLSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestHealthSettingsSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"latency_critical": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"latency_warning": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"packet_loss_critical": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"packet_loss_warning": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"jitter_critical": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"jitter_warning": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"http_latency_critical": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"http_latency_warning": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"http_valid_codes": {
			Type:     schema.TypeList,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"dns_valid_codes": {
			Type:     schema.TypeList,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	})
}

func makeTestMonitoringSettingsSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"activation_grace_period": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"activation_time_unit": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"activation_time_window": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"activation_times": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"notification_channels": {
			Type:     schema.TypeList,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	})
}

func makeTestPingSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"period": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestTraceSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"period": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"protocol": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"port": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"limit": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
	})
}

func makeUserInfoSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"email": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"full_name": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	})
}
