package pinerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mirrorru/dot/pinerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleWrappingError_Produce() {
	var kpe pinerr.WrappingError

	err := func() error {
		// WrappingError remember place of first calling Produce()
		return kpe.Produce(errors.New("incoming_error"))
	}()
	fmt.Println(err)

	// Output:
	// incoming_error @github.com/mirrorru/dot/pinerr_test.ExampleWrappingError_Produce.func1:18
}

func Test_wrappingError_Error(t *testing.T) {
	t.Parallel()
	we := pinerr.WrappingError{}
	srcErr := errors.New("srcErr")
	err := we.Produce(srcErr)
	require.Error(t, err)
	assert.Contains(t, err.Error(), srcErr.Error())
}

func Test_wrappingError_Unwrap(t *testing.T) {
	t.Parallel()
	we := pinerr.WrappingError{}
	srcErr := errors.New("srcErr")
	err := we.Produce(srcErr)
	require.Error(t, err)
	require.ErrorIs(t, err, srcErr)
	uwErr := errors.Unwrap(err)
	require.Error(t, uwErr)
	assert.Equal(t, uwErr, srcErr)
}

func TestWrappingError_Produce(t *testing.T) {
	t.Parallel()
	we := pinerr.WrappingError{}

	srcErr := errors.New("srcErr")
	err1 := we.Produce(srcErr)
	require.Error(t, err1)

	err2 := we.Produce(srcErr)
	require.Error(t, err2)
	assert.Equal(t, err1, err2)

	srcErr3 := errors.New("srcErr3")
	err3 := we.Produce(srcErr3)
	require.Error(t, err3)
	assert.NotEqual(t, err1, err3)
}
