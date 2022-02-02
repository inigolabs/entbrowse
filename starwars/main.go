package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/inigolabs/entbrowse"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed data.yaml
	data []byte
)

func main() {
	ctx := context.Background()

	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	check(err)
	err = client.Schema.Create(ctx)
	check(err)
	ent.LoadYML(ctx, bytes.NewReader(data), client)
	check(err)

	browseIntf := ent.NewBrowseInterface(client)
	handler := entbrowse.Handler(browseIntf)

	fmt.Printf("starwars entbrowse running on port :8080")
	err = http.ListenAndServe(":8080", handler)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
