package entbrowse

import (
	"context"
)

// Browser is implemented by ent to allow for browsing all the
//  entities in the ent schema.
type Browser interface {
	Types() []string
	Count(ctx context.Context, entityType string) int
	List(ctx context.Context, entityType string) interface{}
	Object(ctx context.Context, typeName string, id int64) interface{}
}
