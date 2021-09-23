package synthetics

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	idKey = "id"
)

// schemaMode determines if we want a schema for:
// - reading single item - we need to provide "id" of the item to read, everything else is provided by the server,
// - reading list of items - we don't need to provide a thing, everything is provided by the server,
// - creating new item - we need to provide a bunch of obligatory attributes, the rest is provided by the server.
type schemaMode int

const (
	readSingle schemaMode = iota
	readList
	create
)

func requiredOnCreate(mode schemaMode) bool {
	return mode == create
}

func requiredOnReadSingle(mode schemaMode) bool {
	return mode == readSingle
}

func computedOnCreateAndReadList(mode schemaMode) bool {
	return mode == create || mode == readList
}

func computedOnRead(mode schemaMode) bool {
	return mode == readSingle || mode == readList
}

// makeNestedObjectSchemaRequired returns a list of 1 element to emulate a nested object.
// See: https://learn.hashicorp.com/tutorials/terraform/provider-create?in=terraform/providers#define-order-schema
func makeNestedObjectSchemaRequired(mode schemaMode, properties map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: requiredOnCreate(mode),
		Computed: computedOnRead(mode),
		Elem: &schema.Resource{
			Schema: properties,
		},
	}
}

// makeNestedObjectSchema returns a list of 1 element to emulate a nested object.
// See: https://learn.hashicorp.com/tutorials/terraform/provider-create?in=terraform/providers#define-order-schema
func makeNestedObjectSchemaOptional(mode schemaMode, properties map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: computedOnRead(mode),
		Elem: &schema.Resource{
			Schema: properties,
		},
	}
}

// makeReadOnlyNestedObjectSchema returns a list of 1 element to emulate nested object.
// The object is read-only - only provided by the server.
// See: https://learn.hashicorp.com/tutorials/terraform/provider-create?in=terraform/providers#define-order-schema
func makeReadOnlyNestedObjectSchema(properties map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: properties,
		},
	}
}
