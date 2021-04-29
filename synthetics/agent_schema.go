package synthetics

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

func makeAgentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Type: schema.TypeString,
			Computed: true,
		},
		"status": {
			Type: schema.TypeString,
			Computed: true,
			// enumeration: AGENT_STATUS_UNSPECIFIED, AGENT_STATUS_OK, AGENT_STATUS_WAIT, AGENT_STATUS_DELETED
		},
		"alias": {
			Type: schema.TypeString,
			Computed: true,
		},
		"type": {
			Type: schema.TypeString,
			Computed: true,
		},
		"os": {
			Type: schema.TypeString,
			Computed: true,
		},
		"ip": {
			Type: schema.TypeString,
			Computed: true,
		},
		"lat": {
			Type: schema.TypeFloat,
			Computed: true,
		},
		"long": {
			Type: schema.TypeFloat,
			Computed: true,
		},
		"last_authed": {
			Type: schema.TypeString,
			Computed: true,
		},
		"family": {
			Type: schema.TypeString,
			Computed: true,
			// enumeration: IP_FAMILY_UNSPECIFIED, IP_FAMILY_V4, IP_FAMILY_V6, IP_FAMILY_DUAL
		},
		"asn": {
			Type: schema.TypeInt,
			Computed: true,
		},
		"site_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		"version": {
			Type: schema.TypeString,
			Computed: true,
		},
		"challenge": {
			Type: schema.TypeString,
			Computed: true,
		},
		"city": {
			Type: schema.TypeString,
			Computed: true,
		},
		"region": {
			Type: schema.TypeString,
			Computed: true,
		},
		"country": {
			Type: schema.TypeString,
			Computed: true,
		},
		"test_ids": {
			Type: schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"local_ip": {
			Type: schema.TypeString,
			Computed: true,
		},
	}
}

func agentToMap(obj *synthetics.V202101beta1Agent) map[string]interface{} {
	m := make(map[string]interface{})
	if obj == nil {
		return m
	}

	m["id"] = obj.Id
	m["name"] = obj.Name
	m["status"] = obj.Status
	m["alias"] = obj.Alias
	m["type"] = obj.Type
	m["os"] = obj.Os
	m["ip"] = obj.Ip
	m["lat"] = obj.Lat
	m["long"] = obj.Long
	m["last_authed"] = obj.LastAuthed.Format(time.RFC3339Nano)
	m["family"] = obj.Family
	m["asn"] = obj.Asn
	m["site_id"] = obj.SiteId
	m["version"] = obj.Version
	m["challenge"] = obj.Challenge
	m["city"] = obj.City
	m["region"] = obj.Region
	m["country"] = obj.Country
	m["test_ids"] = obj.TestIds
	m["local_ip"] = obj.LocalIp
	return m
}
