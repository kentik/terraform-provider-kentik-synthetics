package synthetics

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO(dfurman): provide descriptions, when they are specified in the OpenAPI definitions

func makeAgentSchema(mode schemaMode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Computed: computedOnCreateAndReadList(mode),
			Required: requiredOnReadSingle(mode),
		},
		"name": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"status": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"alias": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"os": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"ip": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"lat": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"long": {
			Type:     schema.TypeFloat,
			Computed: computedOnRead(mode),
		},
		"last_authed": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"family": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"asn": {
			Type:     schema.TypeInt,
			Computed: computedOnRead(mode),
		},
		"site_id": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"version": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"challenge": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"city": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"region": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"country": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"test_ids": {
			Type:     schema.TypeList,
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"local_ip": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"cloud_vpc": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
		"agent_impl": {
			Type:     schema.TypeString,
			Computed: computedOnRead(mode),
		},
	}
}
