package entbrowse

import (
	"embed"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

type listEntity struct {
	Limit             int
	SortQuery         string
	List              interface{}
	PageOptions       map[int]bool
	PrevPage          *string
	NextPage          *string
	PageInfo          string
	PaginationEnabled bool
	SortEnabled       bool
}

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
	list := listEntity{List: data}
	w.Header().Add("Content-Type", "text/html")
	_, err := w.Write(htmlHeader)
	check(err)
	err = s.listTemplate.Execute(w, list)
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
	model := listEntity{
		PageOptions: map[int]bool{
			5:   false,
			10:  false,
			50:  false,
			100: false,
			500: false,
		},
		PaginationEnabled: true,
		SortEnabled:       true,
	}
	var limit, offset int
	orderBy := r.URL.Query().Get("orderBy")
	ascSort := r.URL.Query().Get("sortOrder") != "DESC"

	if orderBy != "" {
		model.SortQuery = fmt.Sprintf("&orderBy=%s", orderBy)
		if ascSort {
			model.SortQuery += "&sortOrder=ASC"
		} else {
			model.SortQuery += "&sortOrder=DESC"
		}
	}

	limitQuery, er := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if er == nil {
		limit = int(limitQuery)
	}

	if limit == 0 {
		// default limit
		limit = 100
	}

	model.Limit = limit
	model.PageOptions[limit] = true

	offsetQuery, er := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	if er == nil {
		offset = int(offsetQuery)
	}

	count := s.browser.Count(ctx, entityType)
	model.List = s.browser.List(ctx, entityType, &limit, &offset, toSnakeCase(orderBy), ascSort)
	listSize := sliceLen(model.List)

	if offset > 0 {
		newOffset := int(math.Max(float64(offset-listSize), 0))
		prevPage := fmt.Sprintf("?limit=%d&offset=%d%s", limit, newOffset, model.SortQuery)
		model.PrevPage = &prevPage
	}

	if offset+listSize < count {
		nextPage := fmt.Sprintf("?limit=%d&offset=%d%s", limit, offset+listSize, model.SortQuery)
		model.NextPage = &nextPage
	}

	model.PageInfo = fmt.Sprintf("%d ... %d of %d", offset, offset+listSize, count)

	err = s.listTemplate.Execute(w, model)
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
	"isNotNil":            isNotNil,
	"isSortActive":        isSortActive,
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

func isNotNil(v interface{}) bool {
	return v != nil && !reflect.ValueOf(v).IsNil()
}

func sliceLen(v interface{}) int {
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return reflect.ValueOf(v).Len()
	}
	return 0
}

func isSortActive(sortQuery, orderBy, sortOrder string) bool {
	return sortQuery == fmt.Sprintf("&orderBy=%s&sortOrder=%s", orderBy, sortOrder)
}

// @FIXME - this is tricky, need to find more reliable way to convert fieldName into column name in the DB
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
