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
		{name: "int true", condition: false, onTrue: 1, onFalse: 2, expect: 1},

    {name: "int false", condition: false, onTrue: 11, onFalse: 22, expect: 22},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expect, dot.Iif(tc.condition, tc.onTrue, tc.onFalse))
		})
	}
}
