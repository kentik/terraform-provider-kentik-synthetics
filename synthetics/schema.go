package synthetics

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// schemaMode determines if we want a schema for:
// - reading single item - we need to provide "id" of the item to read, everything else is provided by the server
// - reading list of items - we don't need to provide a thing, everything is provided by the server
// - creating new item - we need to provide a bunch of obligatory attributes, the rest is provided by the server
type schemaMode int

const (
	readSingle schemaMode = iota
	readList
	create
)

func computedOnCreateAndReadList(mode schemaMode) bool {
	return mode == create || mode == readList
}

func computedOnRead(mode schemaMode) bool {
	return mode == readSingle || mode == readList
}

func requiredOnReadSingle(mode schemaMode) bool {
	return mode == readSingle
}

// makeNestedObjectSchema returns a list of 1 element to emulate a nested object.
// See: https://learn.hashicorp.com/tutorials/terraform/provider-create?in=terraform/providers#define-order-schema
func makeNestedObjectSchema(mode schemaMode, properties map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: computedOnRead(mode),
		Elem: &schema.Resource{
			Schema: properties,
		},
	}
}

func formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339Nano)
}
