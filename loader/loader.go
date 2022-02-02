package loader

import (
	"context"
	"io"
	"text/template"

	"entgo.io/ent/entc/gen"

	_ "embed"
)

var (
	//go:embed ent.go.tmpl
	loaderEntTemplate []byte
)

type EntLoader interface {
	LoadYML(ctx context.Context, reader io.Reader) error
	LoadJson(ctx context.Context, reader io.Reader) error
	WriteInput(ctx context.Context, writer io.Writer) error
	DumpInput(ctx context.Context) interface{}
	WriteOutput(ctx context.Context, writer io.Writer) error
	DumpOutput(ctx context.Context) interface{}
	Clear(ctx context.Context)
}

func Template() *gen.Template {
	return gen.MustParse(
		gen.NewTemplate("loader").
			Funcs(loaderTemplateFuncs()).
			Parse(string(loaderEntTemplate)))
}

func loaderTemplateFuncs() template.FuncMap {
	funcs := gen.Funcs
	funcs["hasRequiredEdgeRef"] = hasRequiredEdgeRef
	return funcs
}

func hasRequiredEdgeRef(node *gen.Type) bool {
	for _, edge := range node.Edges {
		if edge.Ref != nil && !edge.Ref.Optional {
			if !edge.Ref.Unique {
				panic("required edge must also be unique")
			}
			return true
		}
	}
	return false
}
