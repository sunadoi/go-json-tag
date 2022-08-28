package tags

import (
	"reflect"
	"strings"
	"unicode"
)

var commonInitialisms = []string{
	"ACL",
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTP",
	"HTTPS",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SQL",
	"SSH",
	"TCP",
	"TLS",
	"TTL",
	"UDP",
	"UI",
	"UID",
	"UUID",
	"URI",
	"URL",
	"UTF8",
	"VM",
	"XML",
	"XMPP",
	"XSRF",
	"XSS",
}

func AddJson[T comparable](target T) any {
	rv := reflect.ValueOf(&target).Elem()
	rt := rv.Type()
	if rt.Kind() == reflect.Ptr {
		rv = reflect.ValueOf(target).Elem()
		rt = rv.Type()
	}

	fields := make([]reflect.StructField, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fields[i] = reflect.StructField{
			Name: field.Name,
			Type: field.Type,
			Tag:  reflect.StructTag(`json:"` + toSnakeCase(field.Name) + `"`),
		}
	}

	t := reflect.StructOf(fields)
	v := reflect.New(t)
	for i := 0; i < rt.NumField(); i++ {
		v.Elem().Field(i).Set(rv.Field(i))
	}

	return v.Interface()
}

func toSnakeCase(s string) string {
	commonInitialismIndexes := make(map[int]bool, len(commonInitialisms))
	for _, word := range commonInitialisms {
		index := strings.Index(s, word)
		if index != -1 {
			for i := index + 1; i < index+len(word); i++ {
				commonInitialismIndexes[i] = true
			}
		}
	}

	b := &strings.Builder{}
	for i, r := range s {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if commonInitialismIndexes[i] {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
