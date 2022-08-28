package tags

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddJson(t *testing.T) {
	t.Parallel()
	type Foo struct {
		ID          int
		Name        string
		TargetURL   string
		HTTPRequest string
	}
	type FooWithJson struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		TargetURL   string `json:"target_url"`
		HTTPRequest string `json:"http_request"`
	}

	type args struct {
		foo Foo
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "add json",
			args: args{foo: Foo{ID: 1, Name: "Name", TargetURL: "TargetURL", HTTPRequest: "HTTPRequest"}},
			want: FooWithJson{ID: 1, Name: "Name", TargetURL: "TargetURL", HTTPRequest: "HTTPRequest"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := AddJson(tt.args.foo)
			var gotB bytes.Buffer
			_ = json.NewEncoder(&gotB).Encode(got)
			var wantB bytes.Buffer
			_ = json.NewEncoder(&wantB).Encode(tt.want)
			if diff := cmp.Diff(gotB.String(), wantB.String()); diff != "" {
				t.Errorf("AddJson() diff = %v", diff)
			}
		})
	}
}

func Test_toSnakeCase(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "convert to snake case",
			args: args{
				s: "HelloWorld",
			},
			want: "hello_world",
		},
		{
			name: "it includes commonInitialisms",
			args: args{
				s: "TargetURL",
			},
			want: "target_url",
		},
		{
			name: "it begins with commonInitialisms",
			args: args{
				s: "HTTPRequest",
			},
			want: "http_request",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := toSnakeCase(tt.args.s); got != tt.want {
				t.Errorf("toSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
