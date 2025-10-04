package dot_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/mirrorru/dot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceToSlice(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		res := dot.SliceToSlice([]int(nil), strconv.Itoa)
		assert.Nil(t, res)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		res := dot.SliceToSlice([]int{}, strconv.Itoa)
		assert.Equal(t, []string{}, res)
	})
	t.Run("not empty", func(t *testing.T) {
		t.Parallel()
		res := dot.SliceToSlice([]int{1, 10, 100}, strconv.Itoa)
		assert.Equal(t, []string{"1", "10", "100"}, res)
	})
}

func TestSliceToSliceError(t *testing.T) {
	t.Parallel()
	errNegative := errors.New("negative input")
	intToString := func(in int) (string, error) {
		if in < 0 {
			return "", errNegative
		}
		return strconv.Itoa(in), nil
	}

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		res, err := dot.SliceToSliceError([]int(nil), intToString)
		require.NoError(t, err)
		assert.Nil(t, res)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		res, err := dot.SliceToSliceError([]int{}, intToString)
		require.NoError(t, err)
		assert.Equal(t, []string{}, res)
	})
	t.Run("not empty", func(t *testing.T) {
		t.Parallel()
		res, err := dot.SliceToSliceError([]int{1, 10, 100}, intToString)
		require.NoError(t, err)
		assert.Equal(t, []string{"1", "10", "100"}, res)
	})
	t.Run("with error", func(t *testing.T) {
		t.Parallel()
		_, err := dot.SliceToSliceError([]int{1, -10, 100}, intToString)
		assert.Error(t, err)
	})
}
