package pinerr_test

import (
	"fmt"
	"testing"

	"github.com/mirrorru/dot/pinerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleStaticError_Produce() {
	var se = pinerr.NewStatic("some_err")

	err := func() error {
		// StaticError remember place of first calling Produce()
		return se.Produce()
	}()
	fmt.Println(err)

	// Output:
	// some_err @github.com/mirrorru/dot/pinerr_test.ExampleStaticError_Produce.func1:17
}

func TestNewStatic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		msg  string
	}{
		{name: "empty message", msg: ""},
		{name: "filled message", msg: "bla-bla"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.NotNil(t, pinerr.NewStatic(tt.msg))
		})
	}
}

func TestProduce(t *testing.T) {
	t.Parallel()

	se := pinerr.NewStatic("bla-bla")
	err1 := se.Produce()
	require.Error(t, err1)
	assert.Contains(t, err1.Error(), "bla-bla")
	assert.Contains(t, err1.Error(), "pinerr_test.TestProduce")

	err2 := se.Produce()
	assert.Equal(t, err1, err2)
}
