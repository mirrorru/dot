package dot_test

import (
	"testing"

	"github.com/mirrorru/dot"
	"github.com/stretchr/testify/assert"
)

func TestIif(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		condition bool
		onTrue    any
		onFalse   any
		expect    any
	}{
		{name: "string true", condition: true, onTrue: "a", onFalse: "b", expect: "a"},
		{name: "string false", condition: false, onTrue: "aa", onFalse: "bb", expect: "bb"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expect, dot.Iif(tc.condition, tc.onTrue, tc.onFalse))
		})
	}
}
