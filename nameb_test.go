package nameb

import (
	"reflect"
	"testing"
)

func TestExtractWords(t *testing.T) {
	tests := []struct {
		src  string
		want []string
	}{
		{"  f o  o  ", []string{"f", "o", "o"}},
		{"f oo", []string{"f", "oo"}},
		{"f.oo", []string{"f", "oo"}},
		{"f-oo", []string{"f", "oo"}},
		{"f_oo", []string{"f", "oo"}},
		{"f_oo", []string{"f", "oo"}},
		{"Foo", []string{"Foo"}},
		{"FooBar", []string{"Foo", "Bar"}},
		{"URL", []string{"URL"}},
		{"HTTPServe", []string{"HTTP", "Serve"}},
		{"Foo2Bar", []string{"Foo2", "Bar"}},
		{"FOO2Bar", []string{"FOO2", "Bar"}},
		{"FOO2bar", []string{"FOO", "2bar"}},
	}

	got := make([]string, 0, 2)

	for _, te := range tests {
		got = got[:0]
		extractWords(te.src, func(word string) {
			got = append(got, word)
		})
		if !reflect.DeepEqual(got, te.want) {
			t.Errorf("extractWords(%q) = %v, want %v", te.src, got, te.want)
		}
	}
}

func TestSnake(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{"", ""},
		{"snake_case", "snake_case"},
		{"kebab-case", "kebab_case"},
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
	}

	for _, te := range tests {
		got := Snake(te.src)
		if got != te.want {
			t.Errorf("Snake(%q) = %q, want %q", te.src, got, te.want)
		}
	}
}

func TestKebab(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{"", ""},
		{"snake_case", "snake-case"},
		{"kebab-case", "kebab-case"},
		{"camelCase", "camel-case"},
		{"PascalCase", "pascal-case"},
	}

	for _, te := range tests {
		got := Kebab(te.src)
		if got != te.want {
			t.Errorf("Kebab(%q) = %q, want %q", te.src, got, te.want)
		}
	}
}

func TestCamel(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{"", ""},
		{"snake_case", "snakeCase"},
		{"kebab-case", "kebabCase"},
		{"camelCase", "camelCase"},
		{"PascalCase", "pascalCase"},
	}

	for _, te := range tests {
		got := Camel(te.src)
		if got != te.want {
			t.Errorf("Camel(%q) = %q, want %q", te.src, got, te.want)
		}
	}
}

func TestPascal(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{"", ""},
		{"snake_case", "SnakeCase"},
		{"kebab-case", "KebabCase"},
		{"camelCase", "CamelCase"},
		{"PascalCase", "PascalCase"},
	}

	for _, te := range tests {
		got := Pascal(te.src)
		if got != te.want {
			t.Errorf("Pascal(%q) = %q, want %q", te.src, got, te.want)
		}
	}
}
