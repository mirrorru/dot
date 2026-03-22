package dot_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mirrorru/dot"
)

func TestNewSet(t *testing.T) {
	t.Parallel()
	const (
		value1 = 10
		value2 = 100
	)
	set := dot.NewSet[int]()
	assert.False(t, set.Contains(value1))
	set.Add(value1)
	assert.False(t, set.Contains(value2))
	assert.True(t, set.Contains(value1))
	set.Add(value1)
	assert.False(t, set.Contains(value2))
	assert.True(t, set.Contains(value1))
	set.Add(value2)
	assert.True(t, set.Contains(value2))
	set.Remove(value1)
	assert.False(t, set.Contains(value1))
	set.Remove(value1)
	assert.False(t, set.Contains(value1))
}
