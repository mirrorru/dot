package repack_test

import (
	"testing"

	"github.com/mirrorru/dot/repack"

	"github.com/stretchr/testify/assert"
)

func TestSplitCamelCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  []string
	}{
		{"", nil},
		{"Any", []string{"Any"}},
		{"AnyKey", []string{"Any", "Key"}},
		{"AnyDBMS", []string{"Any", "DBMS"}},
		{"anyDBMS", []string{"any", "DBMS"}},
		{"anyDBMSKey", []string{"any", "DBMS", "Key"}},
		{"MyLongDBNameForSQL", []string{"My", "Long", "DB", "Name", "For", "SQL"}},
		{"", nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := repack.SplitCamelCase(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"Any", "any"},
		{"OneTwo", "one_two"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := repack.ToSnakeCase(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToKebabCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"Any", "any"},
		{"OneTwo", "one-two"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := repack.ToKebabCase(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
