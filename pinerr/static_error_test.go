package pinerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mirrorru/dot/pinerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleStaticError_Produce() {
	innerError := errors.New("inner error")
	se := pinerr.NewStatic("some_err: %w", innerError)

	err := func() error {
		// StaticError remember place of first calling Produce()
		return se.Produce()
	}()
	fmt.Println(err)
	if errors.Is(err, innerError) {
		fmt.Println("inner error found")
	}

	// Output:
	// some_err: inner error @github.com/mirrorru/dot/pinerr_test.ExampleStaticError_Produce.func1:19
	// inner error found
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

	innerError := errors.New("inner error")
	se := pinerr.NewStatic("some text: %w", innerError)

	err1 := se.Produce()
	require.Error(t, err1)
	assert.Contains(t, err1.Error(), "some text")               //nolint:testifylint
	assert.Contains(t, err1.Error(), "pinerr_test.TestProduce") //nolint:testifylint
	assert.ErrorIs(t, err1, innerError)                         //nolint:testifylint

	err2 := se.Produce()
	assert.Equal(t, err1, err2)
}
