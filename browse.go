package entbrowse

import (
	"embed"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/go-chi/chi"

	jsoniter "github.com/json-iterator/go"
)

var (
	//go:embed header.html
	htmlHeader []byte

	//go:embed footer.html
	htmlFooter []byte

	//go:embed ent.go.tmpl
	browserEntTemplate []byte

	//go:embed entity.go.tmpl
	//go:embed list.go.tmpl
	templateFiles embed.FS

	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Template() *gen.Template {
	return gen.MustParse(
		gen.NewTemplate("admin").
			Funcs(gen.Funcs).
			Parse(string(browserEntTemplate)))
}

// Handler creates an admin interface http handler
func Handler(browser Browser) *chi.Mux {
	s := server{
		browser: browser,
	}

	r := chi.NewRouter()
	r.Get("/favicon.ico", s.faviconHandler)
	r.Get("/", s.rootHandler)
	r.Get("/{type}", s.listHandler)
	r.Get("/{type}/{id}", s.entityHandler)

	s.entityTemplate = template.Must(
		template.New("entity.go.tmpl").Funcs(templateFuncs).
			ParseFS(templateFiles, "entity.go.tmpl"))

	s.listTemplate = template.Must(
		template.New("list.go.tmpl").Funcs(templateFuncs).
			ParseFS(templateFiles, "list.go.tmpl"))

	return r
}

type server struct {
	browser Browser

	entityTemplate *template.Template
	listTemplate   *template.Template
}

func (s *server) faviconHandler(w http.ResponseWriter, r *http.Request) {}

func (s *server) rootHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type Object struct {
		Object string
		Count  int
	}
	var data []Object
	names := s.browser.Types()
	sort.Strings(names)
	for _, name := range names {
		data = append(data, Object{
			Object: linkify(name, "/"+name),
			Count:  s.browser.Count(ctx, name),
		})
	}
	w.Header().Add("Content-Type", "text/html")
	_, err := w.Write(htmlHeader)
	check(err)
	err = s.listTemplate.Execute(w, data)
	check(err)
	_, err = w.Write(htmlFooter)
	check(err)
}

func (s *server) listHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entityType := chi.URLParam(r, "type")
	w.Header().Add("Content-Type", "text/html")
	_, err := w.Write(htmlHeader)
	check(err)
	err = s.listTemplate.Execute(w, s.browser.List(ctx, entityType))
	check(err)
	_, err = w.Write(htmlFooter)
	check(err)
}

func (s *server) entityHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entityType := chi.URLParam(r, "type")
	entityID := s.browser.ConvertID(entityType, chi.URLParam(r, "id"))
	w.Header().Add("Content-Type", "text/html")
	_, err := w.Write(htmlHeader)
	check(err)
	err = s.entityTemplate.Execute(w, s.browser.Object(ctx, entityType, entityID))
	check(err)
	_, err = w.Write(htmlFooter)
	check(err)
}

func linkify(label string, link string) string {
	return fmt.Sprintf("<a href='%s'>%s</a>", link, label)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var templateFuncs = template.FuncMap{
	"pretty":              pretty,
	"isSlice":             isSlice,
	"isNonEmptySlice":     isNonEmptySlice,
	"reflectStructType":   reflectStructType,
	"reflectStructKeys":   reflectStructKeys,
	"reflectStructValues": reflectStructValues,
	"reflectStructItems":  reflectStructItems,
	"reflectHasField":     reflectHasField,
}

func pretty(obj interface{}) string {
	switch o := obj.(type) {
	case string:
		return o
	case int:
		return strconv.Itoa(o)
	}

	objectString, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", obj)
	}
	return string(objectString)
}

func isSlice(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Slice
}

func isNonEmptySlice(obj interface{}) bool {
	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return reflect.ValueOf(obj).Len() > 0
	}
	return false
}

func reflectStructType(obj interface{}) string {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("type must be struct")
	}

	return v.Type().Name()
}

func reflectStructKeys(obj interface{}) []string {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("type must be struct")
	}

	var out []string
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			name := v.Type().Field(i).Name
			if name != "Edges" {
				out = append(out, name)
			}
		}
	}
	return out
}

func reflectStructValues(obj interface{}) []interface{} {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("type must be stuct")
	}
	typeName := reflectStructType(obj)

	var out []interface{}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			name := v.Type().Field(i).Name
			if name == "ID" {
				val := fmt.Sprint(v.Field(i).Interface())
				link := linkify(val, "/"+typeName+"/"+val)
				out = append(out, link)
			} else if name != "Edges" {
				val := v.Field(i).Interface()
				switch valTyped := val.(type) {
				case string:
					if len(valTyped) > 80 {
						val = valTyped[:80] + "..."
					}
				}
				out = append(out, val)
			}
		}
	}

	return out
}

func reflectStructItems(obj interface{}) [][]interface{} {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("type must be struct")
	}

	var out [][]interface{}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			name := v.Type().Field(i).Name
			if name != "Edges" {
				out = append(out, []interface{}{
					v.Type().Field(i).Name,
					v.Field(i).Interface(),
				})
			}
		}
	}
	return out
}

func reflectHasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}
