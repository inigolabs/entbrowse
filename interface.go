package entbrowse

import (
	"context"
)

// Browser is implemented by ent to allow for browsing all the entities in the ent schema.
type Browser interface {
	// Types returns the name of all the ent schema types
	Types() []string

	// ConvertID converts a string id to the native golang type
	ConvertID(typeName string, id string) interface{}

	// Count returns the number of entities of a particular type
	Count(ctx context.Context, entityType string) int

	// List returns a list of entities
	List(ctx context.Context, entityType string, limit, offset *int, sortField string, ascSort bool) interface{}

	// Object returns a particualr entity including all it's edges
	Object(ctx context.Context, typeName string, id interface{}) interface{}
}
