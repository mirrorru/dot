package dot_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/mirrorru/dot"
	"github.com/stretchr/testify/assert"
)

func TestSliceToSlice(t *testing.T) {
	t.Parallel()
	intToString := func(in int) string { return strconv.Itoa(in) }

	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		res := dot.SliceToSlice([]int(nil), intToString)
		assert.Nil(t, res)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		res := dot.SliceToSlice([]int{}, intToString)
		assert.Equal(t, []string{}, res)
	})
	t.Run("not empty", func(t *testing.T) {
		t.Parallel()
		res := dot.SliceToSlice([]int{1, 10, 100}, intToString)
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
		assert.NoError(t, err)
		assert.Nil(t, res)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		res, err := dot.SliceToSliceError([]int{}, intToString)
		assert.NoError(t, err)
		assert.Equal(t, []string{}, res)
	})
	t.Run("not empty", func(t *testing.T) {
		t.Parallel()
		res, err := dot.SliceToSliceError([]int{1, 10, 100}, intToString)
		assert.NoError(t, err)
		assert.Equal(t, []string{"1", "10", "100"}, res)
	})
	t.Run("with error", func(t *testing.T) {
		t.Parallel()
		_, err := dot.SliceToSliceError([]int{1, -10, 100}, intToString)
		assert.Error(t, err)
	})
}
