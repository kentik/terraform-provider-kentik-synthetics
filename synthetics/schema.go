package synthetics

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
